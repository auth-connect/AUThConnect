import { CommonModule } from '@angular/common';
import { Component, inject, OnInit } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { AuthService } from './services/auth-service/auth.service';
import { UserAuthFacade } from './facades/userauth-facade/userauth-facade.service';
import { IUser } from './interfaces/user.interface';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../environment/environment';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {

  title = 'AUTHCONNECT';

  authService = inject(AuthService);
  userAuthFacade = inject(UserAuthFacade);
  http = inject(HttpClient);

  ngOnInit(): void {
    // const headers = new HttpHeaders({
    //   'Content-Type': 'application/json',
    //   Accept: 'application/json',
    //   // Add other custom headers here if necessary
    // });

    // const apiUrl = `${environment.apiUrl}/v1/users/register`;

    // this.http
    //   .get<{ user: IUser }>(apiUrl, { headers })
    //   .subscribe({
    //     next: (response) => {
    //       console.log('response', response);
    //       this.authService.currentUserSig.set(response.user);
    //     },
    //     error: () => {
    //       this.authService.currentUserSig.set(null);
    //     },
    //   });
  }
}
