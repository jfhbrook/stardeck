.PHONY:
sounds: win3x winxp bt floppy/start.mp3 videlectrix/start.mp3

win3x:
	mkdir -p win3x
	curl -L https://winsounds.com/downloads/Windows3x.zip -o win3x/Windows3x.zip
	cd win3x && unzip Windows3x.zip
	rm -f win3x/Windows3x.zip

winxp:
	mkdir -p winxp
	curl -L 'https://archive.org/compress/windowsxpstartup_201910/formats=VBR%20MP3&file=/windowsxpstartup_201910.zip' -o winxp/windowsxp.zip
	cd winxp && unzip windowsxp.zip
	rm -f winxp/windowsxp.zip

bt:
	mkdir -p bt
	curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-is-ready-to-pair.mp3 -o bt/the-bluetooth-device-is-ready-to-pair.mp3
	curl -L https://www.myinstants.com/media/sounds/the-bluetooth-device-its-connected-succesfull.mp3 -o bt/the-bluetooth-device-its-connected-succesfull.mp3

floppy/full.wav:
	mkdir -p floppy
	yt-dlp --extract-audio https://www.youtube.com/watch?v=o_quPha61D0 --audio-format wav --output floppy/full.wav

floppy/start.mp3:
	ffmpeg -y -ss 23 -t 7 -i floppy/full.wav floppy/start.mp3

videlectrix/full.wav:
	mkdir -p videlectrix
	yt-dlp --extract-audio https://www.youtube.com/watch?v=xBmxHT2SUXg --audio-format wav --output videlectrix/full.wav

videlectrix/start.mp3:
	ffmpeg -y -ss 3 -t '5.5' -i videlectrix/full.wav videlectrix/start.mp3
