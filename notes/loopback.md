# loopback

In order to actually play music on the plusdeck, you need to loopback line in to line out.

I have a script that does this. To use it:

```bash
just loopback enable
```

To put it back:

```bash
just loopback disable
```

Here's a useful set of SO posts:

- <https://askubuntu.com/questions/123798/how-to-hear-my-voice-in-speakers-with-a-mic>
- <https://superuser.com/questions/1725016/loopback-listen-to-line-in-alsa-pulseaudio>

Note that the audio coming over line in is **VERY STOUT** - line in needs to be set to about **20%** to get decent audio quality.
