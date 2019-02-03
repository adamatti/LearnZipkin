package adamatti;

import brave.handler.FinishedSpanHandler;
import brave.handler.MutableSpan;
import brave.propagation.TraceContext;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.condition.ConditionalOnMissingBean;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.web.client.RestTemplate;

import static java.util.Objects.isNull;

@Configuration
@SpringBootApplication
public class Main {

	public static void main(String[] args) {
		SpringApplication.run(Main.class, args);
	}

	@Value("${starwars.url}")
	private String starwarsUrl;

	// Required for https
	@Bean
	@ConditionalOnMissingBean
	public RestTemplate restTemplate(RestTemplateBuilder builder) {
		return builder.requestFactory(HttpComponentsClientHttpRequestFactory.class).build();
	}

	// Fix dependency name
	@Bean
	FinishedSpanHandler finishedSpanHandler(){
		return new FinishedSpanHandler() {
			@Override
			public boolean handle(TraceContext context, MutableSpan span) {
				String tagUrl = span.tag("http.url");
				if (!isNull(tagUrl) && tagUrl.startsWith(starwarsUrl)) {
					span.remoteServiceName("star-wars");
				}
				return true;
			}
		};
	}
}

