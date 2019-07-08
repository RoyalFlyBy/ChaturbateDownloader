# ChaturbateDownloader

A CLI Chaturbate.com downloader that allows you to download streams as they are live in the highest quality available without the weird layers on top of the screen.

##### Supports the following
* Download streams without weird overlay (DMCA and other stickers taking up more screen estate than they should -> not anymore)
* Save the stream to your local storage
* Can save for X minutes and automatically quit or save for as long as you want it to run

##### How does it work?
It downloads streams by the link you supply, no need to log in and no encrypted cookies because nothing worthy of protecting is in there.

Depends on FFmpeg (static).

##### Examples

Imagine the link of stream being: ``https://chaturbate.com/mykinkydope/``

Then the command to download indefinitely (until the stream goes private or stops) would be:

``chaturbateDownloader.exe --URL=https://chaturbate.com/mykinkydope/``

However just having the username would work too, as long ass you use the username as it is represented in the chaturbate url (eg: spaces are replaced with underscores):

``chaturbateDownloader.exe --URL=mykinkydope``

The command to download the next 5 minutes of the stream:

``chaturbateDownloader.exe --timeout=5 --URL=https://chaturbate.com/mykinkydope/``

The command that I like to use adds ``[chaturbate.com]`` before the filename

``chaturbateDownloader.exe --withsite=true --URL=https://chaturbate.com/mykinkydope/``

## Installation

First install FFmpeg, on windows this is a little less straightforward than on linux

#### Windows
* Download the STATIC windows package that suits your platform (32 or 64 bit) from https://ffmpeg.org/download.html and then unzip this package.
* When you unzip this package you will find a sub folder called ``bin`` that contains `ffmpeg.exe`
* Now there are 2 options:
    * Take note of the file location of ``ffmpeg.exe``, copy it somewhere as you will need it later
    * Add the contents of the ``bin`` folder to your environment variable called `PATH`
* Download the repository contents
* Run the executable that fits your platform (32 bit: `_386`, 64 bit: `_amd64`)
    * If you added ffmpeg to your environment variables ignore this
    * `chaturbateDownloader.exe --ffmpeg=C:\path\to\ffmpeg.exe --URL=urlhere`
    * Because you did not add ffmpeg to your environment variables this program can't find the ffmpeg executable so be sure to ALWAYS add the ffmpeg flag with the right path to the ffmpeg executab;e

#### Linux

Run your OS's version of:
`sudo apt install ffmpeg`

After that you can run the executable that fits your platform perfectly without the ffmpeg flag

## Disclaimer

Use at own risk.

##### My experience

Nothing much to it, pretty straightforward for now

Unusual stutters are fixed, at the cost of having an external dependency.

##### Future

I don't plan to maintain this unless something breaks due to chaturbate site updates or if a crucial feature is missing

##### Unsupported
* Private streams
* logging in

Why? Unlike the PHDownloader I have no account to test these features with therefore had no way of implementing them

Want them to be implemented? I will need to be able to do some testing, donating money or an account with tokens towards that goal would allow me to implement those cool things.

Simply open up an issue and we will figure it out from there. 
