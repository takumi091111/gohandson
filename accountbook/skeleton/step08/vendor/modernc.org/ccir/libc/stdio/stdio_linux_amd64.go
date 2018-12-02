// Code generated by running "go generate". DO NOT EDIT.

// ----------------------------------------------------------------------------
//      /usr/lib64/gcc/x86_64-suse-linux/4.8/include/stddef.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1989-2013 Free Software Foundation, Inc.

This file is part of GCC.

GCC is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3, or (at your option)
any later version.

GCC is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

Under Section 7 of GPL version 3, you are granted additional
permissions described in the GCC Runtime Library Exception, version
3.1, as published by the Free Software Foundation.

You should have received a copy of the GNU General Public License and
a copy of the GCC Runtime Library Exception along with this program;
see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
<http://www.gnu.org/licenses/>.  */

// ----------------------------------------------------------------------------
//      /usr/include/wchar.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1995-2015 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <http://www.gnu.org/licenses/>.  */

// ----------------------------------------------------------------------------
//      /usr/include/libio.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1991-2015 Free Software Foundation, Inc.
   This file is part of the GNU C Library.
   Written by Per Bothner <bothner@cygnus.com>.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <http://www.gnu.org/licenses/>.

   As a special exception, if you link the code in this file with
   files compiled with a GNU compiler to produce an executable,
   that does not cause the resulting executable to be covered by
   the GNU Lesser General Public License.  This exception does not
   however invalidate any other reasons why the executable file
   might be covered by the GNU Lesser General Public License.
   This exception applies to code released by its copyright holders
   in files containing the exception.  */

// ----------------------------------------------------------------------------
//      /usr/include/bits/stdio_lim.h
// ----------------------------------------------------------------------------
/* Copyright (C) 1994-2015 Free Software Foundation, Inc.
   This file is part of the GNU C Library.

   The GNU C Library is free software; you can redistribute it and/or
   modify it under the terms of the GNU Lesser General Public
   License as published by the Free Software Foundation; either
   version 2.1 of the License, or (at your option) any later version.

   The GNU C Library is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
   Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public
   License along with the GNU C Library; if not, see
   <http://www.gnu.org/licenses/>.  */

package stdio // import "modernc.org/ccir/libc/stdio"

const (
	XBUFSIZ                 = 8192
	XCHAR_BIT               = 8
	XCHAR_MAX               = 127
	XCHAR_MIN               = -128
	XE2BIG                  = 7
	XEACCES                 = 13
	XEADDRINUSE             = 98
	XEADDRNOTAVAIL          = 99
	XEADV                   = 68
	XEAFNOSUPPORT           = 97
	XEAGAIN                 = 11
	XEALREADY               = 114
	XEBADE                  = 52
	XEBADF                  = 9
	XEBADFD                 = 77
	XEBADMSG                = 74
	XEBADR                  = 53
	XEBADRQC                = 56
	XEBADSLT                = 57
	XEBFONT                 = 59
	XEBUSY                  = 16
	XECANCELED              = 125
	XECHILD                 = 10
	XECHRNG                 = 44
	XECOMM                  = 70
	XECONNABORTED           = 103
	XECONNREFUSED           = 111
	XECONNRESET             = 104
	XEDEADLK                = 35
	XEDEADLOCK              = 35
	XEDESTADDRREQ           = 89
	XEDOM                   = 33
	XEDOTDOT                = 73
	XEDQUOT                 = 122
	XEEXIST                 = 17
	XEFAULT                 = 14
	XEFBIG                  = 27
	XEHOSTDOWN              = 112
	XEHOSTUNREACH           = 113
	XEHWPOISON              = 133
	XEIDRM                  = 43
	XEILSEQ                 = 84
	XEINPROGRESS            = 115
	XEINTR                  = 4
	XEINVAL                 = 22
	XEIO                    = 5
	XEISCONN                = 106
	XEISDIR                 = 21
	XEISNAM                 = 120
	XEKEYEXPIRED            = 127
	XEKEYREJECTED           = 129
	XEKEYREVOKED            = 128
	XEL2HLT                 = 51
	XEL2NSYNC               = 45
	XEL3HLT                 = 46
	XEL3RST                 = 47
	XELIBACC                = 79
	XELIBBAD                = 80
	XELIBEXEC               = 83
	XELIBMAX                = 82
	XELIBSCN                = 81
	XELNRNG                 = 48
	XELOOP                  = 40
	XEMEDIUMTYPE            = 124
	XEMFILE                 = 24
	XEMLINK                 = 31
	XEMSGSIZE               = 90
	XEMULTIHOP              = 72
	XENAMETOOLONG           = 36
	XENAVAIL                = 119
	XENETDOWN               = 100
	XENETRESET              = 102
	XENETUNREACH            = 101
	XENFILE                 = 23
	XENOANO                 = 55
	XENOBUFS                = 105
	XENOCSI                 = 50
	XENODATA                = 61
	XENODEV                 = 19
	XENOENT                 = 2
	XENOEXEC                = 8
	XENOKEY                 = 126
	XENOLCK                 = 37
	XENOLINK                = 67
	XENOMEDIUM              = 123
	XENOMEM                 = 12
	XENOMSG                 = 42
	XENONET                 = 64
	XENOPKG                 = 65
	XENOPROTOOPT            = 92
	XENOSPC                 = 28
	XENOSR                  = 63
	XENOSTR                 = 60
	XENOSYS                 = 38
	XENOTBLK                = 15
	XENOTCONN               = 107
	XENOTDIR                = 20
	XENOTEMPTY              = 39
	XENOTNAM                = 118
	XENOTRECOVERABLE        = 131
	XENOTSOCK               = 88
	XENOTSUP                = 95
	XENOTTY                 = 25
	XENOTUNIQ               = 76
	XENXIO                  = 6
	XEOF                    = -1
	XEOPNOTSUPP             = 95
	XEOVERFLOW              = 75
	XEOWNERDEAD             = 130
	XEPERM                  = 1
	XEPFNOSUPPORT           = 96
	XEPIPE                  = 32
	XEPROTO                 = 71
	XEPROTONOSUPPORT        = 93
	XEPROTOTYPE             = 91
	XERANGE                 = 34
	XEREMCHG                = 78
	XEREMOTE                = 66
	XEREMOTEIO              = 121
	XERESTART               = 85
	XERFKILL                = 132
	XEROFS                  = 30
	XESHUTDOWN              = 108
	XESOCKTNOSUPPORT        = 94
	XESPIPE                 = 29
	XESRCH                  = 3
	XESRMNT                 = 69
	XESTALE                 = 116
	XESTRPIPE               = 86
	XETIME                  = 62
	XETIMEDOUT              = 110
	XETOOMANYREFS           = 109
	XETXTBSY                = 26
	XEUCLEAN                = 117
	XEUNATCH                = 49
	XEUSERS                 = 87
	XEWOULDBLOCK            = 11
	XEXDEV                  = 18
	XEXFULL                 = 54
	XFILENAME_MAX           = 4096
	XFOPEN_MAX              = 16
	XINT_MAX                = 2147483647
	XINT_MIN                = 2147483648
	XLLONG_MAX              = 9223372036854775807
	XLLONG_MIN              = -9223372036854775808
	XLONG_MAX               = 9223372036854775807
	XLONG_MIN               = -9223372036854775808
	XL_ctermid              = 9
	XL_cuserid              = 9
	XL_tmpnam               = 20
	XP_tmpdir               = "/tmp"
	XSCHAR_MAX              = 127
	XSCHAR_MIN              = -128
	XSEEK_CUR               = 1
	XSEEK_END               = 2
	XSEEK_SET               = 0
	XSHRT_MAX               = 32767
	XSHRT_MIN               = -32768
	XTMP_MAX                = 238328
	XUCHAR_MAX              = 255
	XUINT_MAX               = 4294967295
	XULLONG_MAX             = 18446744073709551615
	XULONG_MAX              = 18446744073709551615
	XUSHRT_MAX              = 65535
	X_BITS_TYPES_H          = 1
	X_CTYPE_H               = 1
	X_G_BUFSIZ              = 8192
	X_G_HAVE_MMAP           = 1
	X_G_HAVE_MREMAP         = 1
	X_G_HAVE_ST_BLKSIZE     = 0
	X_G_IO_IO_FILE_VERSION  = 131073
	X_G_config_h            = 1
	X_G_va_list             = 0
	X_IOFBF                 = 0
	X_IOLBF                 = 1
	X_IONBF                 = 2
	X_IOS_APPEND            = 8
	X_IOS_ATEND             = 4
	X_IOS_BIN               = 128
	X_IOS_INPUT             = 1
	X_IOS_NOCREATE          = 32
	X_IOS_NOREPLACE         = 64
	X_IOS_OUTPUT            = 2
	X_IOS_TRUNC             = 16
	X_IO_BAD_SEEN           = 16384
	X_IO_BOOLALPHA          = 65536
	X_IO_BUFSIZ             = 8192
	X_IO_CURRENTLY_PUTTING  = 2048
	X_IO_DEC                = 16
	X_IO_DELETE_DONT_CLOSE  = 64
	X_IO_DONT_CLOSE         = 32768
	X_IO_EOF_SEEN           = 16
	X_IO_ERR_SEEN           = 32
	X_IO_FIXED              = 4096
	X_IO_FLAGS2_MMAP        = 1
	X_IO_FLAGS2_NOTCANCEL   = 2
	X_IO_FLAGS2_USER_WBUF   = 8
	X_IO_HAVE_ST_BLKSIZE    = 0
	X_IO_HEX                = 64
	X_IO_INTERNAL           = 8
	X_IO_IN_BACKUP          = 256
	X_IO_IS_APPENDING       = 4096
	X_IO_IS_FILEBUF         = 8192
	X_IO_LEFT               = 2
	X_IO_LINE_BUF           = 512
	X_IO_LINKED             = 128
	X_IO_MAGIC              = 4222418944
	X_IO_MAGIC_MASK         = 4294901760
	X_IO_NO_READS           = 4
	X_IO_NO_WRITES          = 8
	X_IO_OCT                = 32
	X_IO_RIGHT              = 4
	X_IO_SCIENTIFIC         = 2048
	X_IO_SHOWBASE           = 128
	X_IO_SHOWPOINT          = 256
	X_IO_SHOWPOS            = 1024
	X_IO_SKIPWS             = 1
	X_IO_STDIO              = 16384
	X_IO_TIED_PUT_GET       = 1024
	X_IO_UNBUFFERED         = 2
	X_IO_UNIFIED_JUMPTABLES = 1
	X_IO_UNITBUF            = 8192
	X_IO_UPPERCASE          = 512
	X_IO_USER_BUF           = 1
	X_IO_USER_LOCK          = 32768
	X_IO_file_flags         = 0
	X_IO_fpos64_t           = 0
	X_IO_fpos_t             = 0
	X_IO_iconv_t            = 0
	X_IO_off64_t            = 0
	X_IO_off_t              = 0
	X_IO_pid_t              = 0
	X_IO_size_t             = 0
	X_IO_ssize_t            = 0
	X_IO_uid_t              = 0
	X_IO_va_list            = 0
	X_IO_wint_t             = 0
	X_OLD_STDIO_MAGIC       = 4206624768
	X_STDIO_H               = 1
	X__FILE_defined         = 1
	X__S32_TYPE             = 0
	X__SLONG32_TYPE         = 0
	X____FILE_defined       = 1
	X____mbstate_t_defined  = 1
	Xfgetpos                = 0
	Xfopen                  = 0
	Xfreopen                = 0
	Xfscanf                 = 0
	Xfseeko                 = 0
	Xfsetpos                = 0
	Xftello                 = 0
	Xscanf                  = 0
	Xsscanf                 = 0
	Xtmpfile                = 0
	Xvfscanf                = 0
	Xvscanf                 = 0
	Xvsscanf                = 0
	X_ISalnum               = 8
	X_ISalpha               = 1024
	X_ISblank               = 1
	X_IScntrl               = 2
	X_ISdigit               = 2048
	X_ISgraph               = 32768
	X_ISlower               = 512
	X_ISprint               = 16384
	X_ISpunct               = 4
	X_ISspace               = 8192
	X_ISupper               = 256
	X_ISxdigit              = 4096
	X__codecvt_error        = 2
	X__codecvt_noconv       = 3
	X__codecvt_ok           = 0
	X__codecvt_partial      = 1
)
