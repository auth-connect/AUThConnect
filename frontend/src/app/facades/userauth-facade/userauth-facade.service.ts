import { inject, Injectable, signal } from '@angular/core';
import { IUser } from '../../interfaces/user.interface';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { AuthService } from '../../services/auth-service/auth.service';
import { Router } from '@angular/router';
import { environment } from '../../../environment/environment';

import * as UserAuthActions from '../../actions/userauth-actions/userauth.actions';

@Injectable({
  providedIn: 'root',
})
export class UserAuthFacade {
  http = inject(HttpClient);
  router = inject(Router);
  authService = inject(AuthService);

  apiUrl = `${environment.apiUrl}/users`;

  constructor() {}

  registerUser(user: IUser) {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      Accept: 'application/json',
      // Add other custom headers here if necessary
    });

    const apiUrl = `${environment.apiUrl}/v1/users/register`;

    this.http
      .post<{ user: IUser }>(apiUrl, {
        user: user,
      }, { headers })
      .subscribe((response) => {
        console.log('response', response);
        localStorage.setItem('token', response.user.token);
        this.authService.currentUserSig.set(response.user);
        this.router.navigateByUrl('');
      });
    // this.store.dispatch(UserAuthActions.registerUser({ userData: user }));
  }

  // // TODO Create actions
  loginUser(user: IUser) {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json',
      Accept: 'application/json',
      // Add other custom headers here if necessary
    });

    const apiUrl = `${environment.apiUrl}/v1/users/login`;
    this.http
      .post<{ user: IUser }>(apiUrl, {
        user: user,
      }, { headers })
      .subscribe((response) => {
        console.log('response', response);
        localStorage.setItem('token', response.user.token);
        this.authService.currentUserSig.set(response.user);
        this.router.navigateByUrl('');
      });
  }

  // getUser(){
  //     this.http
  //           .get<{ user: UserInterface }>(this.apiUrl)
  //           .subscribe({
  //             next: (response) => {
  //               console.log('response', response);
  //               this.authService.currentUserSig.set(response.user);
  //             },
  //             error: () => {
  //               this.authService.currentUserSig.set(null);
  //             },
  //           });
  // }
}
