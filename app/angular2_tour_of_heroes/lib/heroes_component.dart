import 'package:angular2/core.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'hero_detail_component.dart';
import 'package:angular2_tour_of_heroes/hero_service.dart';
import 'package:angular2/router.dart';
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

    Router _router;

    final HeroService _heroService;
    List<Hero> heroes;
    Hero selectedHero;

    HeroesComponent(this._heroService);

    onSelect(Hero hero) {
        selectedHero = hero;
    }

    getHeroes() async {
        heroes = await _heroService.getHeroes();
    }

    ngOnInit() {
        getHeroes();
    }

    gotoDetail(){
        final link = ['HeroDetail', { 'id': selectedHero}];
        _router.navigate(link);

    }

}