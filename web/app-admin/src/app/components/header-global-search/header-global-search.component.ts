import { Component, signal } from '@angular/core';
import { provideIcons } from '@ng-icons/core';

import {
  radixMagnifyingGlass,
  radixCalendar,
  radixCardStack,
  radixFace,
  radixGear,
  radixPerson,
  radixPlus,
} from '@ng-icons/radix-icons';
@Component({
  selector: 'app-header-global-search',
  standalone: true,
  imports: [],
  providers: [
    provideIcons({
      radixMagnifyingGlass,
      radixCalendar,
      radixCardStack,
      radixFace,
      radixGear,
      radixPerson,
      radixPlus,
    }),
  ],
  templateUrl: './header-global-search.component.html',
  styleUrl: './header-global-search.component.scss',
})
export class HeaderGlobalSearchComponent {}
