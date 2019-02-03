package adamatti;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import java.util.HashMap;
import java.util.Map;


@RestController
@RequestMapping("/")
public class SampleController {
    private Logger log = LoggerFactory.getLogger(this.getClass());

    @Autowired
    private StarWarsRepo starWarsRepo;

    @GetMapping(path="/")
    public ResponseEntity<Map> status(){
        Map map = new HashMap();
        map.put("status","ok");

        return ResponseEntity.ok(map);
    }

    @GetMapping(path="/people/{id}",produces = "application/json")
    public ResponseEntity<Map> findPeople(@PathVariable int id){
        Map people = starWarsRepo.findPeople(id);
        return ResponseEntity.ok(people);
    }
}
