import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-home-listing',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './home-listing.component.html',
  styleUrl: './home-listing.component.scss'
})
export class HomeListingComponent {


  items = [
    {
      name: '',
      image: "",
    },
    {
      name: '',
      image: "",
    },
    {
      name: '',
      image: "",
    },
    {
      name: '',
      image: "",
    },
    {
      name: '',
      image: "",
    },
    {
      name: '',
      image: "",
    },
  ];

  get firstSixItems() {
    return this.items.slice(0, 6).map((item, index) => ({
      ...item,
      index: index
    }));
  }
}
