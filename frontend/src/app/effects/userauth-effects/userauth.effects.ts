import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { mergeMap, map, catchError } from 'rxjs/operators';
import { of } from 'rxjs';
import { UserAuthFacade } from '../../facades/userauth-facade/userauth-facade.service';
import * as UserAuthActions from '../../actions/userauth-actions/userauth.actions';

@Injectable()
export class AuthEffects {
  // registerUser$ = createEffect(() =>
  //   this.actions$.pipe(
  //     ofType(UserAuthActions.registerUser),
  //     mergeMap(action =>
  //       this.userAuthFacade.registerUser(action.userData).pipe(
  //         map(response => UserAuthActions.registerSuccess({ token: response.token })),
  //         catchError(error => of(UserAuthActions.registerFailure({ error })))
  //       )
  //     )
  //   )
  // );


  constructor(private actions$: Actions, private userAuthFacade: UserAuthFacade) {}
}
