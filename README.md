# 

## Install/Setup on Ubuntu server 20.04
## Setup PAM
1. `vim /etc/pam.d/common-session` and add:
```
# and here are more per-package modules (the "Additional" block)
session required        pam_unix.so
session optional        pam_systemd.so
session optional                        pam_mkhomedir.so
# end of pam-auth-update config
```
1. `pam-update`
