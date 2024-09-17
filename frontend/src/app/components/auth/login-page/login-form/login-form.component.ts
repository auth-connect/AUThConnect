import { Component, signal } from '@angular/core';
import { NgIconComponent, provideIcons, provideNgIconsConfig } from '@ng-icons/core';
import { lucideLoader } from '@ng-icons/lucide';

@Component({
  selector: 'app-login-form',
  standalone: true,
  imports: [NgIconComponent],
  viewProviders: [
    provideIcons({ lucideLoader }), 
    provideNgIconsConfig({
      size: '1.7em',
      color: '#d2d3d5',
    }),
  ],
  templateUrl: './login-form.component.html',
  styleUrl: './login-form.component.scss'
})
export class LoginFormComponent {
  isLoading = signal(false);

	send() {
		this.isLoading.set(true);
		setTimeout(() => this.isLoading.set(false), 3000);
	}

}
