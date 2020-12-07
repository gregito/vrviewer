<a href="https://mhssz.hu/"><img src="https://mhssz.hu/wp-content/uploads/2015/07/logo.png" alt="Magyar Hegy- és Sportmászó Szövetség" width="100" height="100"></a> 

## Intro
A - currently - simple application to workaround the sporadic unavailability of the mhssz site's 
<a href="https://vr.mhssz.hu/">competition tracker</a>.

## Build
To compile it to a runnable binary use the make target:

```make build``` or simply  ```make```, since the default target is the build.

Afterwards you can find the compiled binaries under the `/build/Linux|Darwin|Windows/vrviewer`(.exe in case of Windows)

## Usage

Once you grabbed your executable binary from the above-mentioned folder – or from the Releases section of this GitHub 
page – you can use it from your terminal in the following way(s):

```vrviewer "your name"```

If - for some mysterious reason – you would like to have some more information what is going on behind the scenes, set the 
```DEBUG=true``` environment variable in advance to your current command prompt, or execute the application with this 
value inlined before the executable file, like this:

```DEBUG=true vrviewer "your name"```

## Caching

From 0.1.0 the application supports a file based "caching" solution. This means that when the application first fetches
the necessary data it also stores it afterwards in a directory. The reason behind this is that most of the competition 
results are permanent therefore it should not be fetched over and over again which could take a significant amount of 
time based on one's network capacity. Of course, this data may change especially when a new competition would be 
introduced, so over some time we have to re-fetch everything the keep the data kinda fresh.

### Configuring cache temporary directory

One can specify/override the path with a custom one for the fetched data if needed.
For this, you have to set the `VRV_TMPDIR` environment variable. Keep in mind that whatever path you set, it will 
always contain a sub-folder created by the application.

### Configuring cache content renewal interval

At this point the default renewal interval when the content of the cached files will be overwritten is 6 days.
If you would like to specify a longer/shorter interval period, you can do that by setting the `CACHE_RENEWAL_INTERVAL`
environment variable. This field takes the amount of minutes to keep the data.
Please keep in mind that this renewal interval will be applied based on the date/time of the previous fetching
so whether you have restarted the application or your computer, after the given time passed, next time the application
will re-fetch the necessary data.
Also, if you set this value to zero, then the application will use the default time.

## Releases
Sometimes when I feel like to do so I'm "releasing" a seemingly "okayish" version of the current state that you can use 
if you wish. If there is a "published" release that means that version is the most stable and hence preferred for any
kind of usage.

You can find the releases on the right side of the page.

## Unnecessary excuses
At this point, the application has no graphical interface and not yet decided whether it will have or not.
This is only for personal usage since I have no contact with the MHSSZ therefore if some breaking API change happens on 
their side this application may not be operating anymore.

