<a href="https://mhssz.hu/"><img src="https://mhssz.hu/wp-content/uploads/2015/07/logo.png" alt="Magyar Hegy- és Sportmászó Szövetség" width="100" height="100"></a> 

A - currently - simple application to workaround the sporadic unavailability of the mhssz site's <a href="https://vr.mhssz.hu/">competition tracker</a>.

To compile it to a runnable binary use the make target:

```make build``` or simply  ```make```, since the default target is the build.

Afterwards you can find the compiled binaries under the `/build/Linux|Darwin|Windows/` folder with the name of `vrr`(.exe in case of Windows)

At this point the application has no grafical interface and not yet decided whether it will have or not.
This is only for personal usage since I have no contact with the MHSSZ therefore if some breaking API change happens on their site this application may not be operating anymore.
