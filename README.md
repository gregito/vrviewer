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

### Filtering for name(s)
You can search for one or multiple names by passing the names after the `--names` flag separated by colons (but each 
colon should not be followed by a whitespace)

Example for multiple names: ```./vrviewer --names="your name,his name,her name"```

### Filtering for year
If you would like to search for a specific year you can do that by providing the year after the `--year` flag.

e.g.: ```./vrviewer --names="your name" --year=2020```

### Filtering for competition type
When you're interested only in a specific kind of competition type (boulder or lead) you can filter it with the `--type` 
flag.

The valid values are: **BOULDER, LEAD**

e.g.: `./vrviewer --names="your name" --type=BOULDER`

### Filtering for competition name
If you would like to search for a specific competition in a year you can do that by providing the year after 
the `--competition-name` flag.

e.g.: ```./vrviewer --competition-name="Monkey Power 6" --year=2021```

### Listing competitions by year
You can list the competition for a given year by using the `--year` and the `--list-competitions` flags together.
It is necessary and required to define a year.

e.g.: ```./vrviewer --list-competitions --year=2021```

If - for some mysterious reason – you would like to have some more information what is going on behind the scenes, set the 
```DEBUG=1``` environment variable in advance to your current command prompt, or execute the application with this 
value inlined before the executable file, like this:

```DEBUG=1 ./vrviewer --names="your name"```

## Releases
Sometimes when I feel like to do so I'm "releasing" a seemingly "okayish" version of the current state that you can use 
if you wish. If there is a "published" release that means that version is the most stable and hence preferred for any
kind of usage.

You can find the releases on the right side of the page.

## Unnecessary excuses
At this point, the application has no graphical interface and not yet decided whether it will have or not.
This is only for personal usage since I have no contact with the MHSSZ therefore if some breaking API change happens on 
their side this application may not be operating anymore.

