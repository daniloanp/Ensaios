import 'package:angular2/core.dart' show Component, OnInit;
import 'package:angular2_tour_of_heroes/hero_service.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'package:angular2/router.dart';


@Component(
    selector: 'my-dashboard',
    templateUrl: 'dashboard_component.html',
    styleUrls: const ['dashboard_component.css']

)
class DashboardComponent implements OnInit {
//    heroes: Hero[] = [];
//    constructor(private heroService: HeroService) { }
//    ngOnInit() {
//        this.heroService.getHeroes()
//            .then(heroes => this.heroes = heroes.slice(1,5));
//}
//gotoDetail(){ /* not implemented yet */}

    final HeroService _heroService;
    Router _router;

    List<Hero> heroes;

    DashboardComponent(this._router, this._heroService);

    @override
    ngOnInit() async {
        final heroes = await _heroService.getHeroes();
        this.heroes = heroes.sublist(1,5);
    }

    gotoDetail(Hero hero){
            final link = ['HeroDetail', { 'id': hero.id }];
            _router.navigate(link);
        /* not implemented yet */
    }
}