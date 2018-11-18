# MagnetSearch

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Search magnet ~


## Usage

``` shell
NAME:
   magent searcher - Download what you need :)

USAGE:
   magnetSearch [global options] command [command options] [arguments...]

VERSION:
   beta

AUTHOR:
   Juntaran <jacinthmail@gmail.com>

COMMANDS:
     search   MagnetSearch search sth
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --keyword value, -k value  search keyword (default: "sakuramomo")
   --number value, -n value   most return numbers (default: 10)
   --sort value, -s value     sort way: 1-time, 2-hot, 3-size (default: 1)
   --detail, -d               show detail information
   --help, -h                 show help
   --version, -v              print the version
```

## Example

``` shell
magnetSearch search -k dota2 -n 15 -s 3

magnet:?xt=urn:btih:F5A2F4E3CD3CDC53C19BEEA6A2C0EBE22DBEAAE8 	 98.5 GB 	 2017-01-26
magnet:?xt=urn:btih:51CA9448A142DB8FA0B2C5D91033E0C91B4FDDE6 	 34.6 GB 	 2016-02-26
magnet:?xt=urn:btih:A6FF6A8F638AF074F96D8B3FD1553A3827687F3C 	 30.2 GB 	 2017-01-18
magnet:?xt=urn:btih:9078F64500B48A6932C695F57AF203F774341191 	 18.2 GB 	 2015-10-18
magnet:?xt=urn:btih:B28AEB594A8419080911B62B24E0F28887E8B252 	 18.2 GB 	 2015-10-21
magnet:?xt=urn:btih:4EE1E544D4BB05539FEBCDAE1B0377E6FFCD74E4 	 10.6 GB 	 2017-02-22
magnet:?xt=urn:btih:84D967BA752AD244061EEE0B015E59F5928BE568 	 9.6 GB 	 2017-08-09
magnet:?xt=urn:btih:F164448355DBA67AB3F256E786A97011D698322B 	 9.6 GB 	 2017-10-06
magnet:?xt=urn:btih:746B430E396B4B73585CA684C092150F97263768 	 9.1 GB 	 2017-06-12
magnet:?xt=urn:btih:5D94DD2169A7C50AFB16BBDEAA04A1838B70F643 	 9.1 GB 	 2017-06-03
magnet:?xt=urn:btih:BCAE0982B2C520018193848352E6369D92401F18 	 9.1 GB 	 2016-10-07
magnet:?xt=urn:btih:48AAE8566877EC9173ACBC1F02D8D575ACC5DD5C 	 9.1 GB 	 2016-12-12
magnet:?xt=urn:btih:0EAD3C484D124F61C1E6E0E197F5EEA0B747C432 	 9.1 GB 	 2016-11-18
magnet:?xt=urn:btih:3F54BA09C2CF08B5EE6CF9CE6296D0F7582A77C8 	 9.1 GB 	 2016-11-27
magnet:?xt=urn:btih:2104A574CB70C835B1E401A45BD511A956A03382 	 9.0 GB 	 2017-05-30
```

You can get format info by `--detail` or `-d`

``` shell
magnetSearch search -k dota2 -n 15 -s 2 -d

序号: 1
标题: Dota2.rar
磁链: magnet:?xt=urn:btih:31857596596E303E78D1895465F72C5270D10961
大小: 5.2 GB
日期: 2015-10-11
热度: 8621

序号: 2
标题: dota2
磁链: magnet:?xt=urn:btih:B28AEB594A8419080911B62B24E0F28887E8B252
大小: 18.2 GB
日期: 2015-10-21
热度: 1312

序号: 3
标题: DotA2 - Soundtrack
磁链: magnet:?xt=urn:btih:BE13696CE7B3E976AB361CA5406E89C7288E24DB
大小: 140.1 MB
日期: 2016-01-16
热度: 1197

序号: 4
标题: Dota2 Beta (No Crack).rar.torrent
磁链: magnet:?xt=urn:btih:F3ADA6189E150F2B7D18E58817A9A3AB30B3BD37
大小: 14.0 KB
日期: 2015-10-16
热度: 1054

序号: 5
标题: Dota2.iso
磁链: magnet:?xt=urn:btih:A87B0A0E046F0BB373C61B3DCE9ECEDD40D7AA74
大小: 7.0 GB
日期: 2015-10-17
热度: 915

序号: 6
标题: Dota2-27October updated
磁链: magnet:?xt=urn:btih:2ECAAC0A3ABF69B390F782801CAFC863CC549039
大小: 7.2 GB
日期: 2015-10-27
热度: 644

序号: 7
标题: Dota2 Ero Gif (2014-04-15).zip
磁链: magnet:?xt=urn:btih:5B0556F2656B23B11CE62B8EEAD621F0146DF7C3
大小: 364.6 MB
日期: 2015-10-10
热度: 610

序号: 8
标题: Dota2 руссификатор (2).exe
磁链: magnet:?xt=urn:btih:10CC1C1B8F06DD5B918CE52C16D9110E069EFC32
大小: 100.1 MB
日期: 2015-10-24
热度: 431

序号: 9
标题: dota2
磁链: magnet:?xt=urn:btih:AAACAB25B1241CC3E3E737255B9B3902213CB347
大小: 7.9 GB
日期: 2015-11-10
热度: 331

序号: 10
标题: Photoshop Graphics Bundle 8500 -DOTA2
磁链: magnet:?xt=urn:btih:11E1B56D5DF22258065A9B70B4D59C5C4FC67FAA
大小: 5.3 GB
日期: 2016-03-20
热度: 324

序号: 11
标题: DOTA2 Transformation Pack for Windows 7, 8 , 8.1 & 10
磁链: magnet:?xt=urn:btih:81AD28BCDA9E7D33142C83987FE16A9399E7D363
大小: 18.4 MB
日期: 2015-10-15
热度: 315

序号: 12
标题: dota2.iso
磁链: magnet:?xt=urn:btih:3C39A728A0348EEC57074F796F9D45473BDF656D
大小: 4.1 GB
日期: 2015-11-28
热度: 263

序号: 13
标题: Dota2 keygen 2.1 by vladz.rar
磁链: magnet:?xt=urn:btih:0B73D9C9ADF1C76270912F82854372DA997401FF
大小: 1.2 MB
日期: 2015-10-10
热度: 255

序号: 14
标题: DOTA2 StarCh
磁链: magnet:?xt=urn:btih:28C2E949872F8CB6EFADE56D3EB55B134B456D2A
大小: 841.2 MB
日期: 2015-11-19
热度: 235

序号: 15
标题: Dota2.rar
磁链: magnet:?xt=urn:btih:C323FCF2386DA07E783EF13B471200EDED89FFAF
大小: 8.2 GB
日期: 2016-03-19
热度: 223
```

mutil keyword search, you must included by `" "`

``` shell
magnetSearch search -k "dota2 gif" -d -n 10 -s 2

序号: 1
标题: Dota2 Ero Gif (2014-04-15).zip
磁链: magnet:?xt=urn:btih:5B0556F2656B23B11CE62B8EEAD621F0146DF7C3
大小: 364.6 MB
日期: 2015-10-10
热度: 610
```

## License
MIT [©Juntaran](https://github.com/Juntaran)

___
## Reference

- [torrent-cli](https://github.com/chenjiandongx/torrent-cli)