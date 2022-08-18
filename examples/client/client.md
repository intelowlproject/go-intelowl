<!-- Will be revised when I'll add the custom logger and easy ways of setting the client up! -->
# Client
A good client is a client that is easy to use, configurable and customizable to a user’s liking. Hence, the client has 2 great features:
1. Configurable HTTP client
2. Customizable timeouts

## Configurable HTTP client
Now from the documentation, you can see you can pass your `http.Client`. This is to facilitate each user’s requirement and taste! If you don’t pass one (`nil`) a default `http.Client` will be made for you!

## Customizable timeouts
From `IntelOwlClientOptions` you can add your own timeout to your requests as well. 
