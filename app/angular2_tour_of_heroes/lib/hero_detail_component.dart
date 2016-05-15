import 'package:angular2/core.dart';
import 'package:angular2_tour_of_heroes/hero.dart';

@Component(
    selector: 'my-hero-detail',
    template: '''
  <div *ngIf="hero != null">
    <h2>{{hero.name}} details!</h2>
    <div><label>id: </label>{{hero.id}}</div>
    <div>
      <label>name: </label>
      <input [(ngModel)]="hero.name" placeholder="name">
    </div>
  </div>
''')
class HeroDetailComponent {
    @Input()
    Hero hero;
}
