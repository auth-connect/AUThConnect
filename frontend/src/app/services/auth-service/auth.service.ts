import { Injectable, signal } from '@angular/core';
import { IUser } from '../../interfaces/user.interface';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  currentUserSig = signal<IUser | undefined | null>(undefined);
}
