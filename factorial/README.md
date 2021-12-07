# Containerized factorial server

I am picking golang as the language of choice for this exercise: it's succint and has excellent support for HTTP server
functionality.

I am limiting the factorial input data by 20, as anything above that will end up overflowing the basic integer type in golang.
If need be, this can always be extended by going into long arithmetic, for example.

Dockerfile is plain and simple; I'm making it a two-stage build so the sources don't end up in the app image.

Not particularly happy about the error handling in the HTTP handler, but this will suffice for this particular example.
Also, this is go error handling we are talking about ;)

The resulting service will:

- Respond to `GET /` on port `:8080` inside the container.
- Expect a query parameter `number` that has to be an int in range `[0, 20]`.
- Return `400 Bad Request` on:
  - Negative input date
  - Malformed queries
  - Input data greater then 20

Not adding an OpenAPI here, but that would be nice for a larger service.
