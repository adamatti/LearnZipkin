package adamatti;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import java.util.Map;

import static java.util.Objects.isNull;

// https://gitter.im/spring-cloud/spring-cloud-sleuth/archives/2018/06/25
// https://cloud.spring.io/spring-cloud-sleuth/1.3.x/multi/multi__sending_spans_to_zipkin.html
@Component
public class StarWarsRepo {
    @Autowired
    private RestTemplate restTemplate;

    @Value("${starwars.url}")
    private String starwarsUrl;

    //@NewSpan(name = "repo-find-people")
    Map findPeople(int id){
        String url = starwarsUrl + "/people/" + id;
        return restTemplate.getForObject(url,Map.class);
    }
}
