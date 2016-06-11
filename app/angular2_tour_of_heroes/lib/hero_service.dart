import 'package:angular2/core.dart';

import 'mock_heroes.dart';
import 'package:angular2_tour_of_heroes/hero.dart';
import 'dart:async';

@Injectable()
class HeroService {
    Future<List<Hero>> getHeroesSlowly() {
        return new Future.delayed(const Duration(seconds: 2), () => mockHeroes);
    }

    Future<List<Hero>> getHeroes() async => mockHeroes;

    Future<Hero> getHero(int id) async {
         return (await getHeroes()).singleWhere((hero) => hero.id == id);
    }
}