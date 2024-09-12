import { Component } from '@angular/core';
import {
  HlmAccordionContentComponent,
  HlmAccordionDirective,
  HlmAccordionIconDirective,
  HlmAccordionItemDirective,
  HlmAccordionTriggerDirective,
} from '@spartan-ng/ui-accordion-helm';
import { HlmIconComponent } from '../../libs/ui/ui-icon-helm/src';

@Component({
    selector: 'app-about',
    standalone: true,
    imports: [
        HlmAccordionDirective,
        HlmAccordionItemDirective,
        HlmAccordionTriggerDirective,
        HlmAccordionContentComponent,
        HlmAccordionIconDirective,
        HlmIconComponent,
    ],
    template: `
        <div hlmAccordion>
            <div hlmAccordionItem>
                <button hlmAccordionTrigger>
                    Is it accessible?
                    <hlm-icon hlmAccIcon />
                </button>
                <hlm-accordion-content>Yes. It adheres to the WAI-ARIA design pattern.</hlm-accordion-content>
            </div>

            <div hlmAccordionItem>
                <button hlmAccordionTrigger>
                    Is it styled?
                    <hlm-icon hlmAccIcon />
                </button>
                <hlm-accordion-content>
                    Yes. It comes with default styles that match the other components' aesthetics.
                </hlm-accordion-content>
            </div>

            <div hlmAccordionItem>
                <button hlmAccordionTrigger>
                    Is it animated?
                    <hlm-icon hlmAccIcon />
                </button>
                <hlm-accordion-content>
                    Yes. It's animated by default, but you can disable it if you prefer.
                </hlm-accordion-content>
            </div>
        </div>
    `,
})
export class AboutComponent {}