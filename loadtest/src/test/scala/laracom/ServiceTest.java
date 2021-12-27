package laracom;

import java.time.Duration;
import java.util.*;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;
import io.gatling.javaapi.jdbc.*;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;
import static io.gatling.javaapi.jdbc.JdbcDsl.*;

public class ServiceTest extends Simulation {

  {
    HttpProtocolBuilder httpProtocol = http
      .baseUrl("http://192.168.99.103:8080")
      .inferHtmlResources(AllowList(), DenyList(""))
      .acceptHeader("*/*")
      .acceptEncodingHeader("gzip, deflate")
      .contentTypeHeader("application/json")
      .userAgentHeader("PostmanRuntime/7.28.4");
    
    Map<CharSequence, String> headers_0 = new HashMap<>();
    headers_0.put("Postman-Token", "e9e5648c-1589-4610-a4db-8dfb0f8645ee");


    ScenarioBuilder scn = scenario("ServiceTest")
      .exec(
        http("request_0")
          .post("/service/demoService/sayHello")
          .headers(headers_0)
          .body(RawFileBody("laracom/servicetest/0000_request.json"))
      );

	  setUp(scn.injectOpen(atOnceUsers(1))).protocols(httpProtocol);
  }
}
