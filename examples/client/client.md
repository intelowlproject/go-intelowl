<!-- Will be revised when I'll add the custom logger and easy ways of setting the client up! -->
# Client
A good client is a client that is easy to use, configurable and customizable to a user’s liking. Hence, the client has 4 great features:
1. Configurable HTTP client
2. Customizable timeouts
3. Logger
4. Easy ways to create the `IntelOwlClient`

## Configurable HTTP client
Now from the documentation, you can see you can pass your `http.Client`. This is to facilitate each user’s requirement and taste! If you don’t pass one (`nil`) a default `http.Client` will be made for you!

## Customizable timeouts
From `IntelOwlClientOptions` you can add your own timeout to your requests as well.

## Logger
To ease  developers' work go-intelowl provides a logger for easy debugging and tracking! For the logger we used [logrus](https://github.com/sirupsen/logrus) because of 2 reasons:
1. Easy to use
2. Extensible to your liking

## Easy ways to create the `IntelOwlClient`
As you know working with Golang structs is sometimes cumbersome we thought we could provide a simple way to create the client in a way that helps speed up development. This gave birth to the idea of using a `JSON` file to create the IntelOwlClient. The method `NewIntelOwlClientThroughJsonFile` does exactly that. Send the `IntelOwlClientOptions` JSON file path with your http.Client and LoggerParams in this method and you'll get the IntelOwlClient!

