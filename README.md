# ChaturbateDownloader

A CLI Chaturbate.com downloader that allows you to download streams as they are live in the highest quality available without the weird layers on top of the screen.

##### Supports the following
* Download streams without weird overlay (DMCA and other stickers taking up more screen estate than they should -> not anymore)
* Save the stream to your local storage
* Can save for X minutes and automatically quit or save for as long as you want it to run

##### How does it work?
It downloads streams by the link you supply, no need to log in and no encrypted cookies because nothing worthy of protecting is in there.

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

## Disclaimer

Use at own risk.

##### My experience

Nothing much to it, pretty straightforward for now

##### Future

I don't plan to maintain this unless something breaks due to chaturbate site updates or if a crucial feature is missing

##### Unsupported
* Private streams
* logging in

Why? Unlike the PHDownloader I have no account to test these features with therefore had no way of implementing them

Want them to be implemented? I will need to be able to do some testing, donating money or an account with tokens towards that goal would allow me to implement those cool things.

Simply open up an issue and we will figure it out from there. 
