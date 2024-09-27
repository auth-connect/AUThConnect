import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-card',
  standalone: true,
  templateUrl: './card.component.html',
  styleUrls: ['./card.component.scss']
})
export class CardComponent {
  @Input() title!: string;
  @Input() content!: string;
  @Input() imageUrl?: string; // Optional image

  constructor() {}
}
