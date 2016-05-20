import 'package:angular2/core.dart';

import 'mock_heroes.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'dart:async';

@Injectable()
class HeroService {
    Future<List<Hero>> getHeroesSlowly() {
        return new Future.delayed(const Duration(seconds: 0), () => mockHeroes);
    }

    Future<List<Hero>> getHeroes() async => this.getHeroesSlowly();

    Future<Hero> getHero(int id) async {
         return getHeroes().then((heroes) {
            heroes.singleWhere((hero){
                return hero.id == id;
            });
         });
    }
}