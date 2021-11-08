import toobusy from 'toobusy-js'

// Middleware which blocks requests when the Node server is too busy
// now automatically retries the request at another instance of the server if it's too busy
export default (req, res, next) => {
  //toobusy.maxLag(100)
  // toobusy.onLag(function(currentLag) {
  //   console.log("Event loop lag detected! Latency: " + currentLag + "ms");
  // });
  // // Don't send 503s in testing, that's dumb, just wait it out
  if (toobusy()) {
    res.statusCode = 503
    res.end(
      'It looks like API is very busy right now, please try again in a minute.'
    )
  } else {
    next()
  }
}
