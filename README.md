# byoContainerGo

A minimal Linux container runtime built in Go with less than 100 lines.
Based on: https://www.infoq.com/articles/build-a-container-golang/

## Requirements
- Linux kernel >= 5.x
- Root / CAP_SYS_ADMIN privileges
- Go 1.22+

## Build & Run
    go build -o byoContainerGo ./cmd/byocontainer/
    sudo ./byoContainerGo run /bin/bash

## Week 1 — Sprint 1 Review ✓[DONE]

### Tests passing
- TestDefaultFlagsIncludesUTS — PASS
- TestDefaultFlagsIncludesPID — PASS
- TestDefaultFlagsIncludesMNT — PASS
- TestBinaryExistsAndBuilds   — PASS
- TestRunRequiresArguments    — PASS

### Verified working
- UTS namespace isolation (hostname changes don't affect host)
- PID namespace isolation (container runs in own PID namespace)
- MNT namespace isolation (container cannot see host /proc)

## Week 2 — Coming up
- Root filesystem isolation using pivotroot
- Alpine Linux rootfs setup
- Mount namespace fully operational

## Known Limitations
- No networking namespace setup yet
- No filesystem isolation yet (Week 2)
- No cgroups resource limits yet (Week 3)
- Not production safe
