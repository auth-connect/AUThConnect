import { inject, Injectable, signal } from '@angular/core';
import { UserInterface } from '../../interfaces/user.interface';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '../../services/auth-service/auth.service';
import { Router } from '@angular/router';

@Injectable({
    providedIn: 'root',
})
export class AuthFacade {
    http = inject(HttpClient);
    router = inject(Router);
    authService = inject(AuthService);

    // TODO Create actions
    loginUser(user: UserInterface) {
        this.http
            .post<{ user: UserInterface }>(
                'https://api.realworld.io/api/users/login',
                {
                    user: user,
                }
            )
            .subscribe((response) => {
                console.log('response', response);
                localStorage.setItem('token', response.user.token);
                this.authService.currentUserSig.set(response.user);
                this.router.navigateByUrl('');
            });

    }

    getUser(){
        this.http
              .get<{ user: UserInterface }>('https://api.realworld.io/api/user')
              .subscribe({
                next: (response) => {
                  console.log('response', response);
                  this.authService.currentUserSig.set(response.user);
                },
                error: () => {
                  this.authService.currentUserSig.set(null);
                },
              });
    }
}
