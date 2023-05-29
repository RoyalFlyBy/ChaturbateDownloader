# ChaturbateDownloader

A CLI Chaturbate.com downloader that allows you to download streams as they are live in the highest quality available without the weird layers on top of the screen.

### How does it work?

#### Supports the following
* Download streams without weird overlay (DMCA and other stickers taking up more screen estate than they should -> not anymore)
* Save the stream to your local storage
* Can save for X minutes and automatically quit or save for as long as you want it to run

#### Installation
Download the program for your OS [here](https://github.com/RoyalFlyBy/ChaturbateDownloader/releases).

#### Examples

Imagine the link of stream being: ``https://chaturbate.com/mykinkydope/``

Then the command to download indefinitely (until the stream goes private or stops) would be:

``./chaturbateDownloader -url "https://chaturbate.com/mykinkydope/"``

However just having the username would work too, as long ass you use the username as it is represented in the chaturbate url (eg: spaces are replaced with underscores):

``./chaturbateDownloader -url mykinkydope``

The command to download the next 5 minutes of the stream:

``./chaturbateDownloader -cutoff 5 -url "https://chaturbate.com/mykinkydope/"``


You can also add a ```-namefmt``` flag to control the file name like this:

```./chaturbateDownloader -video "https://chaturbate.com/mykinkydope/" -namefmt "[:SITE] :MODEL_NAME at :DATE_DOWNLOAD_STARTED-:TIME_DOWNLOAD_STARTED" ```

Result in:
```[chaturbate.com] mykinkydope at 2023-05-29-02-11-39.ts```


#### For developers
You can access program-friendly downloader updates by adding a ```-daemon``` flag, this will cause the program to print JSON which you could parse and build a UI with.


#### My experience

Nothing much to it, pretty straightforward for now

FFmpeg has been removed as a needed requirement, this means the program is capable of running without any external dependencies.

##### NOTE HOWEVER:
Occasionally the downloaded `.ts` file will be corrupted in the following ways:
* Doesn't play
* Plays all the way but crashes at the end
* Unable to skip around
* No timer showing during playback

Fear not, this issue can be solved very easily with a FFmpeg pass through

Just run this command using your locally installed `ffmpeg` executable:
```shell
ffmpeg -i "filename.ts" -c copy "filename_fixed.ts"
```

This command will remux the file using FFmpeg, fixing those issues all while keeping all the quality.

This inconvenience currently is caused by the muxing library I use, or at the very least how I interact with it. Researching how to fix this...

#### Unsupported
* Private streams
* logging in

Why? I have no account to test these features with therefore had no way of implementing them

Want them to be implemented? I will need to be able to do some testing, donating money or an account with tokens towards that goal would allow me to implement those cool things.

Simply open up an issue and we will figure it out from there.

### Disclaimer

Use at own risk.

### Future

I don't plan to maintain this unless something breaks due to chaturbate site updates or if a crucial feature is missing
