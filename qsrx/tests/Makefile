# Makefile.in generated by automake 1.13.4 from Makefile.am.
# tests/Makefile.  Generated from Makefile.in by configure.

# Copyright (C) 1994-2013 Free Software Foundation, Inc.

# This Makefile.in is free software; the Free Software Foundation
# gives unlimited permission to copy and/or distribute it,
# with or without modifications, as long as this notice is preserved.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY, to the extent permitted by law; without
# even the implied warranty of MERCHANTABILITY or FITNESS FOR A
# PARTICULAR PURPOSE.




am__is_gnu_make = test -n '$(MAKEFILE_LIST)' && test -n '$(MAKELEVEL)'
am__make_running_with_option = \
  case $${target_option-} in \
      ?) ;; \
      *) echo "am__make_running_with_option: internal error: invalid" \
              "target option '$${target_option-}' specified" >&2; \
         exit 1;; \
  esac; \
  has_opt=no; \
  sane_makeflags=$$MAKEFLAGS; \
  if $(am__is_gnu_make); then \
    sane_makeflags=$$MFLAGS; \
  else \
    case $$MAKEFLAGS in \
      *\\[\ \	]*) \
        bs=\\; \
        sane_makeflags=`printf '%s\n' "$$MAKEFLAGS" \
          | sed "s/$$bs$$bs[$$bs $$bs	]*//g"`;; \
    esac; \
  fi; \
  skip_next=no; \
  strip_trailopt () \
  { \
    flg=`printf '%s\n' "$$flg" | sed "s/$$1.*$$//"`; \
  }; \
  for flg in $$sane_makeflags; do \
    test $$skip_next = yes && { skip_next=no; continue; }; \
    case $$flg in \
      *=*|--*) continue;; \
        -*I) strip_trailopt 'I'; skip_next=yes;; \
      -*I?*) strip_trailopt 'I';; \
        -*O) strip_trailopt 'O'; skip_next=yes;; \
      -*O?*) strip_trailopt 'O';; \
        -*l) strip_trailopt 'l'; skip_next=yes;; \
      -*l?*) strip_trailopt 'l';; \
      -[dEDm]) skip_next=yes;; \
      -[JT]) skip_next=yes;; \
    esac; \
    case $$flg in \
      *$$target_option*) has_opt=yes; break;; \
    esac; \
  done; \
  test $$has_opt = yes
am__make_dryrun = (target_option=n; $(am__make_running_with_option))
am__make_keepgoing = (target_option=k; $(am__make_running_with_option))
pkgdatadir = $(datadir)/quaggasrx
pkgincludedir = $(includedir)/quaggasrx
pkglibdir = $(libdir)/quaggasrx
pkglibexecdir = $(libexecdir)/quaggasrx
am__cd = CDPATH="$${ZSH_VERSION+.}$(PATH_SEPARATOR)" && cd
install_sh_DATA = $(install_sh) -c -m 644
install_sh_PROGRAM = $(install_sh) -c
install_sh_SCRIPT = $(install_sh) -c
INSTALL_HEADER = $(INSTALL_DATA)
transform = $(program_transform_name)
NORMAL_INSTALL = :
PRE_INSTALL = :
POST_INSTALL = :
NORMAL_UNINSTALL = :
PRE_UNINSTALL = :
POST_UNINSTALL = :
build_triplet = x86_64-unknown-linux-gnu
host_triplet = x86_64-unknown-linux-gnu
target_triplet = x86_64-unknown-linux-gnu
noinst_PROGRAMS = testsig$(EXEEXT) testbuffer$(EXEEXT) \
	testmemory$(EXEEXT) heavy$(EXEEXT) heavywq$(EXEEXT) \
	heavythread$(EXEEXT) aspathtest$(EXEEXT) testprivs$(EXEEXT) \
	teststream$(EXEEXT) testbgpcap$(EXEEXT) ecommtest$(EXEEXT) \
	testbgpmpattr$(EXEEXT) testchecksum$(EXEEXT) \
	testbgpmpath$(EXEEXT) tabletest$(EXEEXT)
subdir = tests
DIST_COMMON = $(srcdir)/Makefile.in $(srcdir)/Makefile.am \
	$(top_srcdir)/depcomp
ACLOCAL_M4 = $(top_srcdir)/aclocal.m4
am__aclocal_m4_deps = $(top_srcdir)/m4/ax_sys_weak_alias.m4 \
	$(top_srcdir)/m4/libtool.m4 $(top_srcdir)/m4/ltoptions.m4 \
	$(top_srcdir)/m4/ltsugar.m4 $(top_srcdir)/m4/ltversion.m4 \
	$(top_srcdir)/m4/lt~obsolete.m4 \
	$(top_srcdir)/m4/srx_m4_check_arch.m4 \
	$(top_srcdir)/m4/srx_m4_check_libraries.m4 \
	$(top_srcdir)/configure.ac
am__configure_deps = $(am__aclocal_m4_deps) $(CONFIGURE_DEPENDENCIES) \
	$(ACLOCAL_M4)
mkinstalldirs = $(install_sh) -d
CONFIG_HEADER = $(top_builddir)/config.h
CONFIG_CLEAN_FILES =
CONFIG_CLEAN_VPATH_FILES =
PROGRAMS = $(noinst_PROGRAMS)
am_aspathtest_OBJECTS = aspath_test.$(OBJEXT)
aspathtest_OBJECTS = $(am_aspathtest_OBJECTS)
aspathtest_DEPENDENCIES = ../bgpd/libbgp.a ../lib/libzebra.la
AM_V_lt = $(am__v_lt_$(V))
am__v_lt_ = $(am__v_lt_$(AM_DEFAULT_VERBOSITY))
am__v_lt_0 = --silent
am__v_lt_1 = 
am_ecommtest_OBJECTS = ecommunity_test.$(OBJEXT)
ecommtest_OBJECTS = $(am_ecommtest_OBJECTS)
ecommtest_DEPENDENCIES = ../bgpd/libbgp.a ../lib/libzebra.la
am_heavy_OBJECTS = heavy.$(OBJEXT) main.$(OBJEXT)
heavy_OBJECTS = $(am_heavy_OBJECTS)
heavy_DEPENDENCIES = ../lib/libzebra.la
am_heavythread_OBJECTS = heavy-thread.$(OBJEXT) main.$(OBJEXT)
heavythread_OBJECTS = $(am_heavythread_OBJECTS)
heavythread_DEPENDENCIES = ../lib/libzebra.la
am_heavywq_OBJECTS = heavy-wq.$(OBJEXT) main.$(OBJEXT)
heavywq_OBJECTS = $(am_heavywq_OBJECTS)
heavywq_DEPENDENCIES = ../lib/libzebra.la
am_tabletest_OBJECTS = table_test.$(OBJEXT)
tabletest_OBJECTS = $(am_tabletest_OBJECTS)
tabletest_DEPENDENCIES = ../lib/libzebra.la
am_testbgpcap_OBJECTS = bgp_capability_test.$(OBJEXT)
testbgpcap_OBJECTS = $(am_testbgpcap_OBJECTS)
testbgpcap_DEPENDENCIES = ../bgpd/libbgp.a ../lib/libzebra.la
am_testbgpmpath_OBJECTS = bgp_mpath_test.$(OBJEXT)
testbgpmpath_OBJECTS = $(am_testbgpmpath_OBJECTS)
testbgpmpath_DEPENDENCIES = ../bgpd/libbgp.a ../lib/libzebra.la
am_testbgpmpattr_OBJECTS = bgp_mp_attr_test.$(OBJEXT)
testbgpmpattr_OBJECTS = $(am_testbgpmpattr_OBJECTS)
testbgpmpattr_DEPENDENCIES = ../bgpd/libbgp.a ../lib/libzebra.la
am_testbuffer_OBJECTS = test-buffer.$(OBJEXT)
testbuffer_OBJECTS = $(am_testbuffer_OBJECTS)
testbuffer_DEPENDENCIES = ../lib/libzebra.la
am_testchecksum_OBJECTS = test-checksum.$(OBJEXT)
testchecksum_OBJECTS = $(am_testchecksum_OBJECTS)
testchecksum_DEPENDENCIES = ../lib/libzebra.la
am_testmemory_OBJECTS = test-memory.$(OBJEXT)
testmemory_OBJECTS = $(am_testmemory_OBJECTS)
testmemory_DEPENDENCIES = ../lib/libzebra.la
am_testprivs_OBJECTS = test-privs.$(OBJEXT)
testprivs_OBJECTS = $(am_testprivs_OBJECTS)
testprivs_DEPENDENCIES = ../lib/libzebra.la
am_testsig_OBJECTS = test-sig.$(OBJEXT)
testsig_OBJECTS = $(am_testsig_OBJECTS)
testsig_DEPENDENCIES = ../lib/libzebra.la
am_teststream_OBJECTS = test-stream.$(OBJEXT)
teststream_OBJECTS = $(am_teststream_OBJECTS)
teststream_DEPENDENCIES = ../lib/libzebra.la
AM_V_P = $(am__v_P_$(V))
am__v_P_ = $(am__v_P_$(AM_DEFAULT_VERBOSITY))
am__v_P_0 = false
am__v_P_1 = :
AM_V_GEN = $(am__v_GEN_$(V))
am__v_GEN_ = $(am__v_GEN_$(AM_DEFAULT_VERBOSITY))
am__v_GEN_0 = @echo "  GEN     " $@;
am__v_GEN_1 = 
AM_V_at = $(am__v_at_$(V))
am__v_at_ = $(am__v_at_$(AM_DEFAULT_VERBOSITY))
am__v_at_0 = @
am__v_at_1 = 
DEFAULT_INCLUDES = -I. -I$(top_builddir)
depcomp = $(SHELL) $(top_srcdir)/depcomp
am__depfiles_maybe = depfiles
am__mv = mv -f
COMPILE = $(CC) $(DEFS) $(DEFAULT_INCLUDES) $(INCLUDES) $(AM_CPPFLAGS) \
	$(CPPFLAGS) $(AM_CFLAGS) $(CFLAGS)
LTCOMPILE = $(LIBTOOL) $(AM_V_lt) --tag=CC $(AM_LIBTOOLFLAGS) \
	$(LIBTOOLFLAGS) --mode=compile $(CC) $(DEFS) \
	$(DEFAULT_INCLUDES) $(INCLUDES) $(AM_CPPFLAGS) $(CPPFLAGS) \
	$(AM_CFLAGS) $(CFLAGS)
AM_V_CC = $(am__v_CC_$(V))
am__v_CC_ = $(am__v_CC_$(AM_DEFAULT_VERBOSITY))
am__v_CC_0 = @echo "  CC      " $@;
am__v_CC_1 = 
CCLD = $(CC)
LINK = $(LIBTOOL) $(AM_V_lt) --tag=CC $(AM_LIBTOOLFLAGS) \
	$(LIBTOOLFLAGS) --mode=link $(CCLD) $(AM_CFLAGS) $(CFLAGS) \
	$(AM_LDFLAGS) $(LDFLAGS) -o $@
AM_V_CCLD = $(am__v_CCLD_$(V))
am__v_CCLD_ = $(am__v_CCLD_$(AM_DEFAULT_VERBOSITY))
am__v_CCLD_0 = @echo "  CCLD    " $@;
am__v_CCLD_1 = 
SOURCES = $(aspathtest_SOURCES) $(ecommtest_SOURCES) $(heavy_SOURCES) \
	$(heavythread_SOURCES) $(heavywq_SOURCES) $(tabletest_SOURCES) \
	$(testbgpcap_SOURCES) $(testbgpmpath_SOURCES) \
	$(testbgpmpattr_SOURCES) $(testbuffer_SOURCES) \
	$(testchecksum_SOURCES) $(testmemory_SOURCES) \
	$(testprivs_SOURCES) $(testsig_SOURCES) $(teststream_SOURCES)
DIST_SOURCES = $(aspathtest_SOURCES) $(ecommtest_SOURCES) \
	$(heavy_SOURCES) $(heavythread_SOURCES) $(heavywq_SOURCES) \
	$(tabletest_SOURCES) $(testbgpcap_SOURCES) \
	$(testbgpmpath_SOURCES) $(testbgpmpattr_SOURCES) \
	$(testbuffer_SOURCES) $(testchecksum_SOURCES) \
	$(testmemory_SOURCES) $(testprivs_SOURCES) $(testsig_SOURCES) \
	$(teststream_SOURCES)
am__can_run_installinfo = \
  case $$AM_UPDATE_INFO_DIR in \
    n|no|NO) false;; \
    *) (install-info --version) >/dev/null 2>&1;; \
  esac
am__tagged_files = $(HEADERS) $(SOURCES) $(TAGS_FILES) $(LISP)
# Read a list of newline-separated strings from the standard input,
# and print each of them once, without duplicates.  Input order is
# *not* preserved.
am__uniquify_input = $(AWK) '\
  BEGIN { nonempty = 0; } \
  { items[$$0] = 1; nonempty = 1; } \
  END { if (nonempty) { for (i in items) print i; }; } \
'
# Make sure the list of sources is unique.  This is necessary because,
# e.g., the same source file might be shared among _SOURCES variables
# for different programs/libraries.
am__define_uniq_tagged_files = \
  list='$(am__tagged_files)'; \
  unique=`for i in $$list; do \
    if test -f "$$i"; then echo $$i; else echo $(srcdir)/$$i; fi; \
  done | $(am__uniquify_input)`
ETAGS = etags
CTAGS = ctags
DISTFILES = $(DIST_COMMON) $(DIST_SOURCES) $(TEXINFOS) $(EXTRA_DIST)
ACLOCAL = aclocal-1.13
AMTAR = $${TAR-tar}
AM_DEFAULT_VERBOSITY = 1
AR = ar
AUTOCONF = autoconf
AUTOHEADER = autoheader
AUTOMAKE = automake-1.13
AWK = gawk
BABELD = 
BGPD = bgpd
BUILD_TESTS = 
CC = gcc
CCDEPMODE = depmode=gcc3
CFLAGS = -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/include -I/opt/project/srx_test1/srx/../_inst//include -Os -fno-omit-frame-pointer -g -std=gnu99 -Wall -Wsign-compare -Wpointer-arith -Wbad-function-cast -Wwrite-strings -Wmissing-prototypes -Wmissing-declarations -Wchar-subscripts -Wcast-qual
CONFDATE = 20191230
CPP = gcc -E
CPPFLAGS = 
CURSES = 
CYGPATH_W = echo
DEFS = -DHAVE_CONFIG_H $(LOCAL_OPTS) -DSYSCONFDIR=\"$(sysconfdir)/\"
DEPDIR = .deps
DLLTOOL = false
DOC = 
DSYMUTIL = 
DUMPBIN = 
ECHO_C = 
ECHO_N = -n
ECHO_T = 
EGREP = /bin//grep -E
EXEEXT = 
FGREP = /bin//grep -F
GAWK = gawk
GREP = /bin//grep
HAVE_LIBPCREPOSIX = 
HAVE_LIBSRxCryptoAPI = 1
HAVE_LIBSRxProxy = 1
IF_METHOD = if_netlink.o
IF_PROC = 
INCLUDES =  -I.. -I$(top_srcdir) -I$(top_srcdir)/lib -I$(top_builddir)/lib -I$(top_srcdir)/$(SRX_DIR)
INSTALL = /bin//install -c
INSTALL_DATA = ${INSTALL} -m 644
INSTALL_PROGRAM = ${INSTALL}
INSTALL_SCRIPT = ${INSTALL}
INSTALL_STRIP_PROGRAM = $(install_sh) -c -s
IOCTL_METHOD = ioctl.o
IPFORWARD = ipforward_proc.o
ISISD = 
ISIS_TOPOLOGY_DIR = 
ISIS_TOPOLOGY_INCLUDES = 
ISIS_TOPOLOGY_LIB = 
KERNEL_METHOD = 
LATEXMK = /bin/false
LD = /bin//ld -m elf_x86_64
LDFLAGS = -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -L/opt/project/srx_test1/srx/../_inst//lib64/srx 
LIBCAP = 
LIBM = -lm
LIBOBJS = 
LIBPAM = 
LIBREADLINE = 
LIBS = -lcrypt -lSRxProxy -lSRxCryptoAPI   -lrt
LIBTOOL = $(SHELL) $(top_builddir)/libtool
LIB_IPV6 = 
LIB_REGEX = 
LIPO = 
LN_S = ln -s
LTLIBOBJS = 
MAJOR_VER = 4
MAKEINFO = makeinfo
MANIFEST_TOOL = :
MINOR_VER = 2
MKDIR_P = /bin//mkdir -p
MORE_TOOLS = 
MULTIPATH_NUM = 1
NETSNMP_CONFIG = 
NM = /bin//nm -B
NMEDIT = 
OBJDUMP = objdump
OBJEXT = o
OSPF6D = 
OSPFAPI = 
OSPFCLIENT = 
OSPFD = 
OTHER_METHOD = 
OTOOL = 
OTOOL64 = 
PACKAGE = quaggasrx
PACKAGE_BUGREPORT = http://bgpsrx.antd.nist.gov
PACKAGE_NAME = QuaggaSRx
PACKAGE_STRING = QuaggaSRx 0.4.2.7
PACKAGE_TARNAME = quaggasrx
PACKAGE_URL = 
PACKAGE_VERSION = 0.4.2.7
PATH_SEPARATOR = :
PDFLATEX = pdflatex
PERL = /bin//perl
PICFLAGS = -fPIE
PILDFLAGS = -pie
PKG_VER = 0
RANLIB = ranlib
RIPD = ripd
RIPNGD = ripngd
RTREAD_METHOD = rtread_netlink.o
RT_METHOD = rt_netlink.o
SED = sed
SET_MAKE = 
SHELL = /bin/sh
SOLARIS = 
STRIP = strip
UPD_VER = 7
VERSION = 0.4.2.7
VTYSH = 
WATCHQUAGGA = watchquagga
WEAK_ALIAS = attribute
WEAK_ALIAS_CROSSFILE = yes
ZEBRA = zebra
abs_builddir = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx/tests
abs_srcdir = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx/tests
abs_top_builddir = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx
abs_top_srcdir = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx
ac_ct_AR = ar
ac_ct_CC = gcc
ac_ct_DUMPBIN = 
am__include = include
am__leading_dot = .
am__quote = 
am__tar = $${TAR-tar} chof - "$$tardir"
am__untar = $${TAR-tar} xf -
bindir = ${exec_prefix}/bin
build = x86_64-unknown-linux-gnu
build_alias = 
build_cpu = x86_64
build_os = linux-gnu
build_vendor = unknown
builddir = .
datadir = ${datarootdir}
datarootdir = ${prefix}/share
docdir = ${datarootdir}/doc/${PACKAGE_TARNAME}
dvidir = ${docdir}
enable_group = root
enable_user = root
enable_vty_group = 
exampledir = ${prefix}/etc
exec_prefix = ${prefix}
grpc_dir = /opt/project/gobgp_test/gowork/src/srx_grpc
host = x86_64-unknown-linux-gnu
host_alias = 
host_cpu = x86_64
host_os = linux-gnu
host_vendor = unknown
htmldir = ${docdir}
includedir = ${prefix}/include
infodir = ${datarootdir}/info
install_sh = ${SHELL} /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx/install-sh
libdir = ${exec_prefix}/lib64
libexecdir = ${exec_prefix}/libexec
localedir = ${datarootdir}/locale
localstatedir = ${prefix}/var
mandir = ${datarootdir}/man
mkdir_p = $(MKDIR_P)
oldincludedir = /usr/include
pdfdir = ${docdir}
pkgsrcdir = 
pkgsrcrcdir = 
prefix = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx/test_install
program_transform_name = s,x,x,
psdir = ${docdir}
quagga_statedir = /opt/project/gobgp_test/gowork/src/srx_grpc/qsrx/test_install/etc
sbindir = ${exec_prefix}/sbin
sca_dir = /opt/project/srx_test1/srx/../_inst/
sharedstatedir = ${prefix}/com
srcdir = .
srx_dir = /opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install
sysconfdir = ${prefix}/etc
target = x86_64-unknown-linux-gnu
target_alias = 
target_cpu = x86_64
target_os = linux-gnu
target_vendor = unknown
top_build_prefix = ../
top_builddir = ..
top_srcdir = ..
AM_CFLAGS = $(PICFLAGS)
AM_LDFLAGS = $(PILDFLAGS) -ldl
testsig_SOURCES = test-sig.c
testbuffer_SOURCES = test-buffer.c
testmemory_SOURCES = test-memory.c
testprivs_SOURCES = test-privs.c
teststream_SOURCES = test-stream.c
heavy_SOURCES = heavy.c main.c
heavywq_SOURCES = heavy-wq.c main.c
heavythread_SOURCES = heavy-thread.c main.c
aspathtest_SOURCES = aspath_test.c
testbgpcap_SOURCES = bgp_capability_test.c
ecommtest_SOURCES = ecommunity_test.c
testbgpmpattr_SOURCES = bgp_mp_attr_test.c
testchecksum_SOURCES = test-checksum.c
testbgpmpath_SOURCES = bgp_mpath_test.c
tabletest_SOURCES = table_test.c
testsig_LDADD = ../lib/libzebra.la 
testbuffer_LDADD = ../lib/libzebra.la 
testmemory_LDADD = ../lib/libzebra.la 
testprivs_LDADD = ../lib/libzebra.la 
teststream_LDADD = ../lib/libzebra.la 
heavy_LDADD = ../lib/libzebra.la  -lm
heavywq_LDADD = ../lib/libzebra.la  -lm
heavythread_LDADD = ../lib/libzebra.la  -lm
aspathtest_LDADD = ../bgpd/libbgp.a ../lib/libzebra.la $(SRX_CLI_LIB) $(SRX_CRYPTO_API_LIBS) $(SRX_API_LIB)  -lm
testbgpcap_LDADD = ../bgpd/libbgp.a ../lib/libzebra.la $(SRX_CLI_LIB) $(SRX_CRYPTO_API_LIBS) $(SRX_API_LIB)  -lm
ecommtest_LDADD = ../bgpd/libbgp.a ../lib/libzebra.la $(SRX_CLI_LIB) $(SRX_CRYPTO_API_LIBS) $(SRX_API_LIB)  -lm
testbgpmpattr_LDADD = ../bgpd/libbgp.a ../lib/libzebra.la $(SRX_CLI_LIB) $(SRX_CRYPTO_API_LIBS) $(SRX_API_LIB)  -lm
testchecksum_LDADD = ../lib/libzebra.la 
testbgpmpath_LDADD = ../bgpd/libbgp.a ../lib/libzebra.la $(SRX_CLI_LIB) $(SRX_CRYPTO_API_LIBS) $(SRX_API_LIB)  -lm
tabletest_LDADD = ../lib/libzebra.la  -lm
all: all-am

.SUFFIXES:
.SUFFIXES: .c .lo .o .obj
$(srcdir)/Makefile.in:  $(srcdir)/Makefile.am  $(am__configure_deps)
	@for dep in $?; do \
	  case '$(am__configure_deps)' in \
	    *$$dep*) \
	      ( cd $(top_builddir) && $(MAKE) $(AM_MAKEFLAGS) am--refresh ) \
	        && { if test -f $@; then exit 0; else break; fi; }; \
	      exit 1;; \
	  esac; \
	done; \
	echo ' cd $(top_srcdir) && $(AUTOMAKE) --gnu tests/Makefile'; \
	$(am__cd) $(top_srcdir) && \
	  $(AUTOMAKE) --gnu tests/Makefile
.PRECIOUS: Makefile
Makefile: $(srcdir)/Makefile.in $(top_builddir)/config.status
	@case '$?' in \
	  *config.status*) \
	    cd $(top_builddir) && $(MAKE) $(AM_MAKEFLAGS) am--refresh;; \
	  *) \
	    echo ' cd $(top_builddir) && $(SHELL) ./config.status $(subdir)/$@ $(am__depfiles_maybe)'; \
	    cd $(top_builddir) && $(SHELL) ./config.status $(subdir)/$@ $(am__depfiles_maybe);; \
	esac;

$(top_builddir)/config.status: $(top_srcdir)/configure $(CONFIG_STATUS_DEPENDENCIES)
	cd $(top_builddir) && $(MAKE) $(AM_MAKEFLAGS) am--refresh

$(top_srcdir)/configure:  $(am__configure_deps)
	cd $(top_builddir) && $(MAKE) $(AM_MAKEFLAGS) am--refresh
$(ACLOCAL_M4):  $(am__aclocal_m4_deps)
	cd $(top_builddir) && $(MAKE) $(AM_MAKEFLAGS) am--refresh
$(am__aclocal_m4_deps):

clean-noinstPROGRAMS:
	@list='$(noinst_PROGRAMS)'; test -n "$$list" || exit 0; \
	echo " rm -f" $$list; \
	rm -f $$list || exit $$?; \
	test -n "$(EXEEXT)" || exit 0; \
	list=`for p in $$list; do echo "$$p"; done | sed 's/$(EXEEXT)$$//'`; \
	echo " rm -f" $$list; \
	rm -f $$list

aspathtest$(EXEEXT): $(aspathtest_OBJECTS) $(aspathtest_DEPENDENCIES) $(EXTRA_aspathtest_DEPENDENCIES) 
	@rm -f aspathtest$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(aspathtest_OBJECTS) $(aspathtest_LDADD) $(LIBS)

ecommtest$(EXEEXT): $(ecommtest_OBJECTS) $(ecommtest_DEPENDENCIES) $(EXTRA_ecommtest_DEPENDENCIES) 
	@rm -f ecommtest$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(ecommtest_OBJECTS) $(ecommtest_LDADD) $(LIBS)

heavy$(EXEEXT): $(heavy_OBJECTS) $(heavy_DEPENDENCIES) $(EXTRA_heavy_DEPENDENCIES) 
	@rm -f heavy$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(heavy_OBJECTS) $(heavy_LDADD) $(LIBS)

heavythread$(EXEEXT): $(heavythread_OBJECTS) $(heavythread_DEPENDENCIES) $(EXTRA_heavythread_DEPENDENCIES) 
	@rm -f heavythread$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(heavythread_OBJECTS) $(heavythread_LDADD) $(LIBS)

heavywq$(EXEEXT): $(heavywq_OBJECTS) $(heavywq_DEPENDENCIES) $(EXTRA_heavywq_DEPENDENCIES) 
	@rm -f heavywq$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(heavywq_OBJECTS) $(heavywq_LDADD) $(LIBS)

tabletest$(EXEEXT): $(tabletest_OBJECTS) $(tabletest_DEPENDENCIES) $(EXTRA_tabletest_DEPENDENCIES) 
	@rm -f tabletest$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(tabletest_OBJECTS) $(tabletest_LDADD) $(LIBS)

testbgpcap$(EXEEXT): $(testbgpcap_OBJECTS) $(testbgpcap_DEPENDENCIES) $(EXTRA_testbgpcap_DEPENDENCIES) 
	@rm -f testbgpcap$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testbgpcap_OBJECTS) $(testbgpcap_LDADD) $(LIBS)

testbgpmpath$(EXEEXT): $(testbgpmpath_OBJECTS) $(testbgpmpath_DEPENDENCIES) $(EXTRA_testbgpmpath_DEPENDENCIES) 
	@rm -f testbgpmpath$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testbgpmpath_OBJECTS) $(testbgpmpath_LDADD) $(LIBS)

testbgpmpattr$(EXEEXT): $(testbgpmpattr_OBJECTS) $(testbgpmpattr_DEPENDENCIES) $(EXTRA_testbgpmpattr_DEPENDENCIES) 
	@rm -f testbgpmpattr$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testbgpmpattr_OBJECTS) $(testbgpmpattr_LDADD) $(LIBS)

testbuffer$(EXEEXT): $(testbuffer_OBJECTS) $(testbuffer_DEPENDENCIES) $(EXTRA_testbuffer_DEPENDENCIES) 
	@rm -f testbuffer$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testbuffer_OBJECTS) $(testbuffer_LDADD) $(LIBS)

testchecksum$(EXEEXT): $(testchecksum_OBJECTS) $(testchecksum_DEPENDENCIES) $(EXTRA_testchecksum_DEPENDENCIES) 
	@rm -f testchecksum$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testchecksum_OBJECTS) $(testchecksum_LDADD) $(LIBS)

testmemory$(EXEEXT): $(testmemory_OBJECTS) $(testmemory_DEPENDENCIES) $(EXTRA_testmemory_DEPENDENCIES) 
	@rm -f testmemory$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testmemory_OBJECTS) $(testmemory_LDADD) $(LIBS)

testprivs$(EXEEXT): $(testprivs_OBJECTS) $(testprivs_DEPENDENCIES) $(EXTRA_testprivs_DEPENDENCIES) 
	@rm -f testprivs$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testprivs_OBJECTS) $(testprivs_LDADD) $(LIBS)

testsig$(EXEEXT): $(testsig_OBJECTS) $(testsig_DEPENDENCIES) $(EXTRA_testsig_DEPENDENCIES) 
	@rm -f testsig$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(testsig_OBJECTS) $(testsig_LDADD) $(LIBS)

teststream$(EXEEXT): $(teststream_OBJECTS) $(teststream_DEPENDENCIES) $(EXTRA_teststream_DEPENDENCIES) 
	@rm -f teststream$(EXEEXT)
	$(AM_V_CCLD)$(LINK) $(teststream_OBJECTS) $(teststream_LDADD) $(LIBS)

mostlyclean-compile:
	-rm -f *.$(OBJEXT)

distclean-compile:
	-rm -f *.tab.c

include ./$(DEPDIR)/aspath_test.Po
include ./$(DEPDIR)/bgp_capability_test.Po
include ./$(DEPDIR)/bgp_mp_attr_test.Po
include ./$(DEPDIR)/bgp_mpath_test.Po
include ./$(DEPDIR)/ecommunity_test.Po
include ./$(DEPDIR)/heavy-thread.Po
include ./$(DEPDIR)/heavy-wq.Po
include ./$(DEPDIR)/heavy.Po
include ./$(DEPDIR)/main.Po
include ./$(DEPDIR)/table_test.Po
include ./$(DEPDIR)/test-buffer.Po
include ./$(DEPDIR)/test-checksum.Po
include ./$(DEPDIR)/test-memory.Po
include ./$(DEPDIR)/test-privs.Po
include ./$(DEPDIR)/test-sig.Po
include ./$(DEPDIR)/test-stream.Po

.c.o:
	$(AM_V_CC)$(COMPILE) -MT $@ -MD -MP -MF $(DEPDIR)/$*.Tpo -c -o $@ $<
	$(AM_V_at)$(am__mv) $(DEPDIR)/$*.Tpo $(DEPDIR)/$*.Po
#	$(AM_V_CC)source='$<' object='$@' libtool=no \
#	DEPDIR=$(DEPDIR) $(CCDEPMODE) $(depcomp) \
#	$(AM_V_CC_no)$(COMPILE) -c $<

.c.obj:
	$(AM_V_CC)$(COMPILE) -MT $@ -MD -MP -MF $(DEPDIR)/$*.Tpo -c -o $@ `$(CYGPATH_W) '$<'`
	$(AM_V_at)$(am__mv) $(DEPDIR)/$*.Tpo $(DEPDIR)/$*.Po
#	$(AM_V_CC)source='$<' object='$@' libtool=no \
#	DEPDIR=$(DEPDIR) $(CCDEPMODE) $(depcomp) \
#	$(AM_V_CC_no)$(COMPILE) -c `$(CYGPATH_W) '$<'`

.c.lo:
	$(AM_V_CC)$(LTCOMPILE) -MT $@ -MD -MP -MF $(DEPDIR)/$*.Tpo -c -o $@ $<
	$(AM_V_at)$(am__mv) $(DEPDIR)/$*.Tpo $(DEPDIR)/$*.Plo
#	$(AM_V_CC)source='$<' object='$@' libtool=yes \
#	DEPDIR=$(DEPDIR) $(CCDEPMODE) $(depcomp) \
#	$(AM_V_CC_no)$(LTCOMPILE) -c -o $@ $<

mostlyclean-libtool:
	-rm -f *.lo

clean-libtool:
	-rm -rf .libs _libs

ID: $(am__tagged_files)
	$(am__define_uniq_tagged_files); mkid -fID $$unique
tags: tags-am
TAGS: tags

tags-am: $(TAGS_DEPENDENCIES) $(am__tagged_files)
	set x; \
	here=`pwd`; \
	$(am__define_uniq_tagged_files); \
	shift; \
	if test -z "$(ETAGS_ARGS)$$*$$unique"; then :; else \
	  test -n "$$unique" || unique=$$empty_fix; \
	  if test $$# -gt 0; then \
	    $(ETAGS) $(ETAGSFLAGS) $(AM_ETAGSFLAGS) $(ETAGS_ARGS) \
	      "$$@" $$unique; \
	  else \
	    $(ETAGS) $(ETAGSFLAGS) $(AM_ETAGSFLAGS) $(ETAGS_ARGS) \
	      $$unique; \
	  fi; \
	fi
ctags: ctags-am

CTAGS: ctags
ctags-am: $(TAGS_DEPENDENCIES) $(am__tagged_files)
	$(am__define_uniq_tagged_files); \
	test -z "$(CTAGS_ARGS)$$unique" \
	  || $(CTAGS) $(CTAGSFLAGS) $(AM_CTAGSFLAGS) $(CTAGS_ARGS) \
	     $$unique

GTAGS:
	here=`$(am__cd) $(top_builddir) && pwd` \
	  && $(am__cd) $(top_srcdir) \
	  && gtags -i $(GTAGS_ARGS) "$$here"
cscopelist: cscopelist-am

cscopelist-am: $(am__tagged_files)
	list='$(am__tagged_files)'; \
	case "$(srcdir)" in \
	  [\\/]* | ?:[\\/]*) sdir="$(srcdir)" ;; \
	  *) sdir=$(subdir)/$(srcdir) ;; \
	esac; \
	for i in $$list; do \
	  if test -f "$$i"; then \
	    echo "$(subdir)/$$i"; \
	  else \
	    echo "$$sdir/$$i"; \
	  fi; \
	done >> $(top_builddir)/cscope.files

distclean-tags:
	-rm -f TAGS ID GTAGS GRTAGS GSYMS GPATH tags

distdir: $(DISTFILES)
	@srcdirstrip=`echo "$(srcdir)" | sed 's/[].[^$$\\*]/\\\\&/g'`; \
	topsrcdirstrip=`echo "$(top_srcdir)" | sed 's/[].[^$$\\*]/\\\\&/g'`; \
	list='$(DISTFILES)'; \
	  dist_files=`for file in $$list; do echo $$file; done | \
	  sed -e "s|^$$srcdirstrip/||;t" \
	      -e "s|^$$topsrcdirstrip/|$(top_builddir)/|;t"`; \
	case $$dist_files in \
	  */*) $(MKDIR_P) `echo "$$dist_files" | \
			   sed '/\//!d;s|^|$(distdir)/|;s,/[^/]*$$,,' | \
			   sort -u` ;; \
	esac; \
	for file in $$dist_files; do \
	  if test -f $$file || test -d $$file; then d=.; else d=$(srcdir); fi; \
	  if test -d $$d/$$file; then \
	    dir=`echo "/$$file" | sed -e 's,/[^/]*$$,,'`; \
	    if test -d "$(distdir)/$$file"; then \
	      find "$(distdir)/$$file" -type d ! -perm -700 -exec chmod u+rwx {} \;; \
	    fi; \
	    if test -d $(srcdir)/$$file && test $$d != $(srcdir); then \
	      cp -fpR $(srcdir)/$$file "$(distdir)$$dir" || exit 1; \
	      find "$(distdir)/$$file" -type d ! -perm -700 -exec chmod u+rwx {} \;; \
	    fi; \
	    cp -fpR $$d/$$file "$(distdir)$$dir" || exit 1; \
	  else \
	    test -f "$(distdir)/$$file" \
	    || cp -p $$d/$$file "$(distdir)/$$file" \
	    || exit 1; \
	  fi; \
	done
check-am: all-am
check: check-am
all-am: Makefile $(PROGRAMS)
installdirs:
install: install-am
install-exec: install-exec-am
install-data: install-data-am
uninstall: uninstall-am

install-am: all-am
	@$(MAKE) $(AM_MAKEFLAGS) install-exec-am install-data-am

installcheck: installcheck-am
install-strip:
	if test -z '$(STRIP)'; then \
	  $(MAKE) $(AM_MAKEFLAGS) INSTALL_PROGRAM="$(INSTALL_STRIP_PROGRAM)" \
	    install_sh_PROGRAM="$(INSTALL_STRIP_PROGRAM)" INSTALL_STRIP_FLAG=-s \
	      install; \
	else \
	  $(MAKE) $(AM_MAKEFLAGS) INSTALL_PROGRAM="$(INSTALL_STRIP_PROGRAM)" \
	    install_sh_PROGRAM="$(INSTALL_STRIP_PROGRAM)" INSTALL_STRIP_FLAG=-s \
	    "INSTALL_PROGRAM_ENV=STRIPPROG='$(STRIP)'" install; \
	fi
mostlyclean-generic:

clean-generic:

distclean-generic:
	-test -z "$(CONFIG_CLEAN_FILES)" || rm -f $(CONFIG_CLEAN_FILES)
	-test . = "$(srcdir)" || test -z "$(CONFIG_CLEAN_VPATH_FILES)" || rm -f $(CONFIG_CLEAN_VPATH_FILES)

maintainer-clean-generic:
	@echo "This command is intended for maintainers to use"
	@echo "it deletes files that may require special tools to rebuild."
clean: clean-am

clean-am: clean-generic clean-libtool clean-noinstPROGRAMS \
	mostlyclean-am

distclean: distclean-am
	-rm -rf ./$(DEPDIR)
	-rm -f Makefile
distclean-am: clean-am distclean-compile distclean-generic \
	distclean-tags

dvi: dvi-am

dvi-am:

html: html-am

html-am:

info: info-am

info-am:

install-data-am:

install-dvi: install-dvi-am

install-dvi-am:

install-exec-am:

install-html: install-html-am

install-html-am:

install-info: install-info-am

install-info-am:

install-man:

install-pdf: install-pdf-am

install-pdf-am:

install-ps: install-ps-am

install-ps-am:

installcheck-am:

maintainer-clean: maintainer-clean-am
	-rm -rf ./$(DEPDIR)
	-rm -f Makefile
maintainer-clean-am: distclean-am maintainer-clean-generic

mostlyclean: mostlyclean-am

mostlyclean-am: mostlyclean-compile mostlyclean-generic \
	mostlyclean-libtool

pdf: pdf-am

pdf-am:

ps: ps-am

ps-am:

uninstall-am:

.MAKE: install-am install-strip

.PHONY: CTAGS GTAGS TAGS all all-am check check-am clean clean-generic \
	clean-libtool clean-noinstPROGRAMS cscopelist-am ctags \
	ctags-am distclean distclean-compile distclean-generic \
	distclean-libtool distclean-tags distdir dvi dvi-am html \
	html-am info info-am install install-am install-data \
	install-data-am install-dvi install-dvi-am install-exec \
	install-exec-am install-html install-html-am install-info \
	install-info-am install-man install-pdf install-pdf-am \
	install-ps install-ps-am install-strip installcheck \
	installcheck-am installdirs maintainer-clean \
	maintainer-clean-generic mostlyclean mostlyclean-compile \
	mostlyclean-generic mostlyclean-libtool pdf pdf-am ps ps-am \
	tags tags-am uninstall uninstall-am


# Tell versions [3.59,3.63) of GNU make to not export all variables.
# Otherwise a system limit (for SysV at least) may be exceeded.
.NOEXPORT:
