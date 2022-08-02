# Endpoints
Now go-intelowl was made with simplicity, ease of use in mind while using Go's best features to make a roboust client library. The client library uses Intelowl's REST API to easily communicate with your intelowl instance. For accessibilty and flexibility we divided the API endpoints as `services` or `resources` where we group together endpoints that preform related operations. For example:
1. `GET /api/tags`
2. `GET /api/tags/{id}`
3. `POST /api/tags/`
4. `PUT /api/tags/{id}`
5. `DELETE /api/tags/{id}`

As you can see these endpoints fall in the `tag` service. Therefore, we make a service struct for `tag` and the `IntelOwlClient` uses that service object to access the tag endpoints.

Now if you're bored with theoretical detail you can see the [example](./endpoints.go) :)

Now some of the endpoints have optional parameters to get a good overview and some running examples you can see it [here](../optionalParams/optionalParams.md)

# References
- [Writing a good Go client](https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c)