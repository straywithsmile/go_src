package poll

import (
	"os"
	"syscall"
)

type pollster struct {
	kq		int
	eventbuf	[4096]syscall.Kevent_t
	kbuf		[1]syscall.Kevent_t
}

func NewPoll() (p *pollster, err error) {
	p = new(pollster)
	if p.kq, err = syscall.Kqueue(); err != nil {
		return nil, os.NewSyscallError("kqueue", err)
	}
	syscall.CloseOnExec(p.kq)
	return p, nil
}

func (p *pollster) AddFD(fd int, mode int, repeat bool) (bool, error) {
	// pollServer is locked.

	var kmode int
	if mode == 'r' {
		kmode = syscall.EVFILT_READ
	} else {
		kmode = syscall.EVFILT_WRITE
	}
	ev := &p.kbuf[0]
	// EV_ADD - add event to kqueue list
	// EV_ONESHOT - delete the event the first time it triggers
	flags := syscall.EV_ADD
	if !repeat {
		flags |= syscall.EV_ONESHOT
	}
	syscall.SetKevent(ev, fd, kmode, flags)

	n, err := syscall.Kevent(p.kq, p.kbuf[:], nil, nil)
	if err != nil {
		return false, os.NewSyscallError("kevent", err)
	}
	if n != 1 || (ev.Flags&syscall.EV_ERROR) == 0 || int(ev.Ident) != fd || int(ev.Filter) != kmode {
		return false, os.NewSyscallError("kqueue phase error", err)
	}
	if ev.Data != 0 {
		return false, syscall.Errno(int(ev.Data))
	}
	return false, nil
}

func (p *pollster) DelFD(fd int, mode int) bool {
	// pollServer is locked.

	var kmode int
	if mode == 'r' {
		kmode = syscall.EVFILT_READ
	} else {
		kmode = syscall.EVFILT_WRITE
	}
	ev := &p.kbuf[0]
	// EV_DELETE - delete event from kqueue list
	syscall.SetKevent(ev, fd, kmode, syscall.EV_DELETE)
	syscall.Kevent(p.kq, p.kbuf[:], nil, nil)
	return false
}

type net_event struct {
	fd	int
	mode	int
	data	int
}

func (p *pollster) Kevent(nsec int64) (result int, events []net_event, err error) {
	var t *syscall.Timespec
	if nsec > 0 {
		if t == nil {
			t = new(syscall.Timespec)
		}
		*t = syscall.NsecToTimespec(nsec)
	}
	n, err := syscall.Kevent(p.kq, nil, p.eventbuf[:], t)

	if err != nil {
		if err == syscall.EINTR {
			return 0, nil, nil
		} else {
			return -1, nil, os.NewSyscallError("kevent", err)
		}
	}
	if n == 0 {
		return -2, nil, nil
	}
	events = make([]net_event, n)
	for i, ev := range p.eventbuf {
		events[i].fd = int(ev.Ident)
		events[i].data = int(ev.Data)
		if ev.Filter == syscall.EVFILT_READ {
			events[i].mode = 'r'
		} else {
			events[i].mode = 'w'
		}
	}
	return n, events, nil
}
