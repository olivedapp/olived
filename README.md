# [Olived](https://olived.app)

[![Latest Release](https://img.shields.io/github/v/release/olivedapp/olived?style=for-the-badge)](https://github.com/olivedapp/olived/releases/latest)
[![Docker Pulls](https://img.shields.io/docker/pulls/olivedapp/olived-web.svg?style=for-the-badge)](https://hub.docker.com/r/olivedapp/olived-web)
[![Docs](https://img.shields.io/badge/DOCS-View-green.svg?style=for-the-badge)](https://olived.app/docs/sites)
[![Changelog](https://img.shields.io/badge/CHANGELOG-What's_New-purple.svg?style=for-the-badge)](https://olived.app/docs/changelog)

## Intro

Olived is a modern live stream recorder that automatically captures your favorite streams when they go live.

Simple, automatic, and built for the streams you never want to miss.

Your favorite streams. Saved.

> ⭐ Star this repository to stay up to date with new releases and improvements.

## Installation

### Desktop

Download the latest release from [⁠GitHub Releases](https://github.com/olivedapp/olived/releases).

#### macOS

macOS may prevent Olived from running because the application is downloaded from the internet. If this happens, remove the quarantine attribute before launching Olived:

```sh
sudo xattr -r -d com.apple.quarantine /Applications/OlivedPro.app
```

### Web

Download the latest release from [⁠GitHub Releases](https://github.com/olivedapp/olived/releases).
Look for the release file containing web in its filename.
Once the executable is running, open your browser and visit: [http://127.0.0.1:9843/](http://127.0.0.1:9843/)

You can customize the port and password using the following options:

```sh
--port 9843 --password yourpassword
```

#### macOS

If macOS prevents the application from running, remove the quarantine attribute:

```sh
sudo xattr -r -d com.apple.quarantine /path/to/OlivedPro_web
```

### Docker

#### Docker Compose

Use the [compose.olived.yaml](https://github.com/olivedapp/olived/blob/main/compose.olived.yaml) file in the project root to start Olived with Docker Compose.

```sh
docker compose up -d
```

Once started, the Web interface is available at:

```sh
http://<your-server-ip>:9843
```

#### Docker Run

You can also run Olived directly with Docker:

```sh
docker run -d \
  --name olived \
  --restart unless-stopped \
  --pull always \
  -p 9843:9843 \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Etc/UTC \
  -e PORT=9843 \
  -e PASSWORD=yourpassword \
  -v $PWD/olivedpro_downloads:/olivedapp/olivedpro_downloads \
  -v $PWD/.olivedpro:/olivedapp/.olivedpro \
  olivedapp/olived-web:latest
```

## Supported Sites

<!-- supported-sites:start -->
| Site ID | Site Name | Live URL Format |
| --- | --- | --- |
| afreecatv | Afreecatv | https://play.afreecatv.com/sol3712/247185916 |
| beeshow | Beeshow | https://www.beeshow.tv/251234653 |
| bigo | BIGO | https://www.bigo.tv/J8023<br>https://slink.bigovideo.tv/GwkE2e |
| bilibili | Bilibili | https://live.bilibili.com/1796100546 |
| bongacams | BongaCams | https://cn.bongacams.com/emmaheart (NSFW) |
| cam4 | CAM4 | https://cam4.com/nickifrenchy (NSFW) |
| cams | CAMS | https://cams.com/Milena_Lynn (NSFW) |
| camsoda | Camsoda | https://www.camsoda.com/isa-luxury (NSFW) |
| catshow168 | Catshow168 | https://h.catshow168.com/register/ShareLive?anchorUid=19583338 |
| chaturbate | Chaturbate | https://chaturbate.com/tina__kim (NSFW) |
| dlive | DLive | https://www.dlive.tv |
| douyin | DOUYIN | https://live.douyin.com/969280804278 |
| douyu | DOUYU | https://www.douyu.com/3168536 |
| dreamcam | Dreamcam | https://dreamcam.com/live/SilverrMoon (NSFW) |
| fansly | Fansly | https://fansly.com/titsmintsalad (NSFW) |
| flextv | Flextv | https://www.flextv.co.kr/channels/376056/live |
| flirt4free | FLIRT4FREE | https://www.flirt4free.com/?model=bonnie-bow (NSFW) |
| haixiutv | Haixiutv | https://www.haixiutv.com/6096233 |
| huajiao | Huajiao | https://www.huajiao.com/user/269174777 |
| huya | HUYA | https://www.huya.com/11420 |
| imkktv | IMKKTV | https://www.imkktv.com/h5/share/video.html?roomId=1785803 |
| immomo | IMMOMO | https://live-api.immomo.com/fep/momo/fe-live-projects/mobile-share/index.html?roomid=17854055882517 |
| inke | INKE | https://mlive9.inke.cn/app/hot/live?uid=762293234 |
| jd | JD | https://zhibo.jd.com/liveroom?liveId=47436664 |
| kick | KICK | https://kick.com/cg5natch |
| kktv5 | KK | https://www.kktv5.com/info/139870658 |
| kuaishou | Kuaishou | https://live.kuaishou.com/u/ktoookt_ |
| lailianjie | Lailianjie | https://show.lailianjie.com/10001816<br>https://show.336769.com/10102851 |
| lang | LANG | https://www.lang.live/room/5259232 |
| lehaitv | Lehaitv | https://www.lehaitv.com/8057460 |
| liveme | LiveMe | https://www.liveme.com/livehot/streaming/78001392<br>https://www.liveme.com/au/m/v/17862088176256942272/index.html |
| livesg | LIVESG | https://h.livesg.cn/live/preview.html?uid=19124369&anchorUid=19106954 |
| massivecams | MASSIVECAMS | https://massivecams.com/cam/Hara99 (NSFW) |
| micous | Mico | https://www.micous.com/live/1104478194 |
| micoworld | Mico | https://www.micoworld.net/live/1104478194 |
| missevan | Missevan | https://fm.missevan.com/live/869083853 |
| mixch | ミクチャ | https://mixch.tv/u/18084988/live |
| myfreecams | MyFreeCams | https://www.myfreecams.com/#amberbeloved (NSFW) |
| naver | 치지직 | https://chzzk.naver.com/live/09f2c3d669fb6ef362f2d78029fc4300 |
| nimo | NIMO | https://www.nimo.tv/live/5602072741 |
| pandalive | PandaTV | https://www.pandalive.co.kr/live/play/cmcm5646 |
| picarto | Picarto | https://picarto.tv/RokusyoKoku |
| pixiv | Pixiv | https://sketch.pixiv.net/@out-mode |
| popkontv | popkontv | https://www.popkontv.com/live/view?castId=88588&partnerCode=P-00001 |
| pornhublive | PornHubLive | https://pornhublive.com/cam/LLspinz (NSFW) |
| rengzu | RENGZU | https://www.rengzu.com/800826185 |
| sexlikereal | SEXLIKEREAL | https://www.sexlikereal.com/vr-cam-girls/bonnie_wright (NSFW) |
| shopee | Shopee | https://live.shopee.co.id/521180499 |
| showroom-live | ショールーム | https://www.showroom-live.com/r/NOA_ROOM |
| showself | Showself | https://www.showself.com/7112709 |
| sooplive | SOOP | https://play.sooplive.co.kr/rlatldgus/277793432 |
| spankbanglive | SpankBangLIVE | https://spankbanglive.com/DeliciaBagnoli (NSFW) |
| streamate | Streamate | https://streamate.com/cam/LeaThomsonn (NSFW) |
| stripchat | STRIPCHAT | https://stripchat.com/Remy_Larson (NSFW) |
| taobao | TAOBAO | https://www.taobao.com |
| tiktok | TikTok | https://www.tiktok.com/@realrammyy/live |
| tlclw | TLCLW | https://www.tlclw.com/101320 |
| ttinglive | 띵라이브 | https://www.ttinglive.com/channels/472369/live |
| twitcasting | Twitcast | https://twitcasting.tv/mille_tomo |
| twitch | Twitch | https://www.twitch.tv/et_1231 |
| vvxqiu | VVXQIU | https://www.vvxqiu.com |
| weibo | WEIBO | https://weibo.com/l/wblive/p/show/1022:2321324933968775742003 |
| winktv | Winktv | https://www.winktv.co.kr/live/play/sasa10046 |
| xhamsterlive | XHAMSTERLIVE | https://xhamsterlive.com/Maddieblairxo (NSFW) |
| xiaohongshu | Xiaohongshu | https://www.xiaohongshu.com/hina/livestream/569318810940914205<br>https://xhslink.com/zfknEQ |
| xlovecam | XLOVECAM | https://xlovecam.com/Milena_Lynn (NSFW) |
| ybw1666 | Ybw1666 | https://www.ybw1666.com/800002967 |
| youtube | YouTube | https://www.youtube.com/@LofiGirl |
| yy | yy | https://www.yy.com/54880976/54880976 |
| 163 | 163 | https://cc.163.com/6761057 |
| 17 | 17LIVE | https://17.live/live/723 |
| 173 | 173 | https://www.173.com/11 |
| 224545 | LUMI | https://www.224545.com/30002438 |
| 6 | 6 | https://v.6.cn/91211 |
| 660807 | 660807 | https://live.660807.com/801566729 |
| 9xiu | 9XIU | https://www.9xiu.com/117991129<br>https://www.9xiuzb.com/124179189 |
<!-- supported-sites:end -->

Note: This list is not exhaustive. Support for additional platforms is continuously being added.

## Feautres

- Automatic Recording: Automatically starts recording when your favorite streamer goes live.
- Simultaneous Recordings: Record multiple live streams at the same time.
- High-Quality Recording: Capture live streams in crisp, high-quality video, up to 4K.
- Automatic MP4 Conversion: Automatically convert recorded files to MP4 for easy playback and editing.
- Deep Customization: Fine-tune recording settings to match your needs.
- Simple by Design: A clean, intuitive interface that makes recording effortless.
- Broad Format Support: Supports HLS, FLV, and other popular streaming formats.
- Smart Watchlist: Keep track of your favorite streamers and automatically record when they go live.
- Reliable Link Processing: Handles complex stream URLs reliably for stable recording.
- Live Notifications: Get notified when your favorite streams go live or change status.

## License

See the [LICENSE](https://olived.app/legal/license) for details.
