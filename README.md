# Mailing Server


## Description

This server acts like a form service, serves as backend for multiple static sites.

Accepts POST requests from allowed static sites and sends email to the specified email address based on the URL configured to my websites.

> confidential information is stored in the .env file.
> 
> Contains the following keys:
> - GOOGLE_CLIENT_ID
> - GOOGLE_CLIENT_SECRET
> - GOOGLE_REFRESH_TOKEN

use the /google_login and /google_callback routes to get the refresh token.

