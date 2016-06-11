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
    List<Hero> heroes;
    final HeroService _heroService;
    final Router _router;


    DashboardComponent(this._heroService, this._router); // Router comming before service


    ngOnInit() async {
        final heroes = await _heroService.getHeroes();
        this.heroes = heroes.sublist(1,5);//Todo check lite, maybe use.toList().
    }

    gotoDetail(Hero hero){
            var link = ['HeroDetail', {'id': hero.id.toString()}];
            _router.navigate(link);
    }
}