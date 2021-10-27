#!/bin/sh
#
# ripd is part of the quagga routing beast
#
# PROVIDE: ripd
# REQUIRE: zebra
##

PATH=/sbin:/bin:/usr/sbin:/usr/bin:/opt/project/gobgp_test/gowork/src/srx_grpc/quagga-srx-v6/../_inst/sbin:/opt/project/gobgp_test/gowork/src/srx_grpc/quagga-srx-v6/../_inst/bin
export PATH

if [ -f /etc/rc.subr ]
then
	. /etc/rc.subr
fi

name="ripd"
rcvar=$name
required_files="${prefix}/etc/${name}.conf"
command="/opt/project/gobgp_test/gowork/src/srx_grpc/quagga-srx-v6/../_inst/sbin/${name}"
command_args="-d"

start_precmd="zebra_precmd"
socket_dir=${prefix}/var
pidfile="${socket_dir}/${name}.pid"

zebra_precmd()
{
    rc_flags="$(
	set -- $rc_flags
	while [ $# -ne 0 ]; do
	    if [ X"$1" = X-P -o X"$1" = X-A ]; then
		break
	    fi
	    shift
	done
	if [ $# -eq 0 ]; then
	    echo "-P 0"
	fi
	) $rc_flags"
}

load_rc_config $name
run_rc_command "$1"