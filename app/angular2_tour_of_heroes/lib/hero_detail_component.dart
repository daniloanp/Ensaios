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
    Hero hero;
    final HeroService _heroService;
    final RouteParams _routeParams;

    HeroDetailComponent(this._heroService, this._routeParams);

    ngOnInit() async {
        var id = int.parse(_routeParams.get('id'));
        this.hero = await (_heroService.getHero(id));
    }

    goBack() {
        window.history.back();
    }
}
