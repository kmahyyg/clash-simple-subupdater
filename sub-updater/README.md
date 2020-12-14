# Phicomm-K2P-Clash

Subscribe updater.

No TUN! No Game!

# Note

To use this tool, you must put all binaries under the same folder.

All config should be put to the path you've specified.

And also clash-dashboard should be downloaded and extracted to corresponding folder.

PLEASE DO NOT DISABLE DNSMASQ DNS ROLE.
THIS SOFTWARE HAS A FEATURE TO DETECT INTERNET SO THAT IT COULD DOWNLOAD REMOTE CONFIG.


# Compile

```bash
$ CGO_ENABLED=0 GOARCH=mipsle GOOS=linux GOMIPS=softfloat go build -trimpath -ldflags '-s -w' -o main ./main.go
$ upx ./main -o ./main.upx
$ upx --version
upx 4.0.0-git-3781df9da238
UCL data compression library 1.03
zlib data compression library 1.2.11
LZMA SDK version 4.43
Copyright (C) 1996-2020 Markus Franz Xaver Johannes Oberhumer
Copyright (C) 1996-2020 Laszlo Molnar
Copyright (C) 2000-2020 John F. Reiser
Copyright (C) 2002-2020 Jens Medoch
Copyright (C) 1995-2005 Jean-loup Gailly and Mark Adler
Copyright (C) 1999-2006 Igor Pavlov
UPX comes with ABSOLUTELY NO WARRANTY; for details type 'upx -L'.
```

# Thanks to

- https://github.com/Dreamacro/clash/wiki/configuration
- https://opt.cn2qq.com/
- https://github.com/vernesong/OpenClash
- https://lancellc.gitbook.io/clash
- https://zyfdegh.github.io/post/202002-go-compile-for-mips/
- https://github.com/upx/upx/tree/devel

# License


 sub-updater
 Copyright (C) 2020  kmahyyg @ PatMeow Ltd.
 
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.
 
 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

