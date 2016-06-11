import 'package:angular2/core.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'hero_detail_component.dart';
import 'package:angular2_tour_of_heroes/hero_service.dart';
import 'package:angular2/router.dart';
import 'dart:async';
// #docregion hero-class-1


// #enddocregion hero-class-1
@Component(
    selector: 'my-heroes',
    templateUrl: 'heroes_component.html',
    styleUrls: const ['heroes_component.css'],
    directives: const [
        HeroDetailComponent
    ]
//    providers: const [HeroService]
)
class HeroesComponent implements OnInit {
    final Router _router;
    final HeroService _heroService;
    List<Hero> heroes;
    Hero selectedHero;

    HeroesComponent(this._heroService, this._router);

    onSelect(Hero hero) {
        selectedHero = hero;
    }

    Future getHeroes() async {
        heroes = await _heroService.getHeroes();
    }

    void ngOnInit() {
        getHeroes();
    }

    Future gotoDetail(){
        final link = ['HeroDetail', { 'id': selectedHero.id.toString()}];
        _router.navigate(link);
    }

}