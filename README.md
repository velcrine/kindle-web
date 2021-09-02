# kindle-web

    // Bit of background
    Amazon Kindle is proprietary/non-OSS. So are it's file extensions like .aax and all.
    Hence, the source of free surfing is the "experimental" web browser, which is based on firefox.

Kindle is good for eyes. We Devs need it for reading thousands of lines of documentations.
I personally utilize it for reading envoy & kubernetes documentation.

The height of hypocrisy of Amazon is such that, they have blocked even playing audio from web.
One such proof comes from the audio-player.html attached. Kindle just block audio player from it.

# Rendering web-pages to fit in kindle

There is "Article Mode". It fits the screen properly, at the cost of hiding code segments, 
high-CSSed content and bubbled information/key-notes.

This app recreates the article mode with no content loss! 

Cost is, you need to download them, write a css based ones in example dir(envoy, kubernetes),
run 

    ./kindle-web -dir envoy
and load them in documents directory in the Kindle reader.