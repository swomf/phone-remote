# phone-remote

xdg-desktop-portal-hyprland doesn't work with KDE Connect,
but I want to sent the key press down button from my phone.
So here we are.

## usage

```bash
go build
sudo ./phone-remote
```

Then access the webpage on port 8000.

## is this secure?

no it's literally a post request asking to press an allowed uinput button
