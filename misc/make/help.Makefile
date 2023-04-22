dep-gawk:
	@ if [ -z "$(shell command -v gawk)" ]; then  \
		if [ -x /usr/local/bin/brew ]; then $(MAKE) _brew_gawk_install; exit 0; fi; \
		if [ -x /usr/bin/apt-get ]; then $(MAKE) _ubuntu_gawk_install; exit 0; fi; \
		if [ -x /usr/bin/yum ]; then  $(MAKE) _centos_gawk_install; exit 0; fi; \
		if [ -x /sbin/apk ]; then  $(MAKE) _alpine_gawk_install; exit 0; fi; \
		echo  "GNU Awk Required, We cannot determine your OS or Package manager. Please install it yourself.";\
		exit 1; \
	fi

_brew_gawk_install:
	@ echo "Instaling gawk using brew... "
	@ brew install gawk --quiet
	@ echo "done"

_ubuntu_gawk_install:
	@ echo "Instaling gawk using apt-get... "
	@ apt-get -q install gawk -y
	@ echo "done"

_alpine_gawk_install:
	@ echo "Instaling gawk using yum... "
	@ apk add --update --no-cache gawk
	@ echo "done"

_centos_gawk_install:
	@ echo "Instaling gawk using yum... "
	@ yum install -q -y gawk;
	@ echo "done"

help: dep-gawk
	@cat $(MAKEFILE_LIST) | \
		grep -E '^# ~~~ .*? [~]+$$|^[a-zA-Z0-9_-]+:.*?## .*$$' | \
		awk '{if ( $$1=="#" ) { \
			match($$0, /^# ~~~ (.+?) [~]+$$/, a);\
			{print "\n", a[1], ""}\
		} else { \
			match($$0, /^([0-9a-zA-Z_-]+):.*?## (.*)$$/, a); \
			{printf "  - \033[32m%-20s\033[0m %s\n",   a[1], a[2]} \
 		}}'
	@echo ""