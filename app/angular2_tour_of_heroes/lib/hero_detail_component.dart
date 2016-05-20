import 'package:angular2/core.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'package:angular2_tour_of_heroes/hero_service.dart';
import 'package:angular2/router.dart';
import 'dart:html' show window;

@Component(
    selector: 'my-hero-detail',
    templateUrl: 'hero_detail_component.html',
    styleUrls: const ['hero_detail_component.css']
)
class HeroDetailComponent implements OnInit {
    HeroService _heroService;
    RouteParams _routeParams;

    HeroDetailComponent(this._heroService, this._routeParams);
    Hero hero;

    @override
    ngOnInit() async {
        String sid = _routeParams.get('id');
//        int id = int.parse(sid+'');
//        window.alert("oi0");
////        print(id);
//        window.alert("oi1");
//        this.hero = await _heroService.getHero(id);
        this.hero = await _heroService.getHero(11);
        window.alert("oi2");
    }

    goBack() {
        window.history.back();
    }
}
