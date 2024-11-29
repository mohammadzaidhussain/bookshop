<template>
    <ion-button :id="id" :class="styleClass" @click="onButtonClick" :disabled="isDisabled">
        {{ title }}
        <ion-spinner :name="spinnerName" v-if="isDisabled"></ion-spinner>
    </ion-button>
</template>

<script>
import { IonButton, IonSpinner } from '@ionic/vue';

export default {
    name: 'AppButton',
    components: {
        IonButton,
        IonSpinner
    },
    props: {
        id: {
            type: String,
            default : ''
        },
        title: {
            type: String,
        },
        styleClass: {
            type: String,
            default: ''
        },
        spinnerName: {
            type: String,
            default: 'bubbles'
        },
        action: {
            type: Function,
            Required: true
        },
        dependentIdsToBeBlocked: {
            type: Array,
            default: []
        }
    },
    data() {
        return {
            isButtonDisabled: false
        }
    },
    computed: {
        isDisabled() {
            return this.isButtonDisabled
        }
    },
    methods: {
        async onButtonClick() {
            this.isButtonDisabled = true;
            const t = setTimeout(() => {
                this.isButtonDisabled = false;
            }, 5000);
            await this.action(this.dependentIdsToBeBlocked);
            clearTimeout(t);
            this.isButtonDisabled = false;
        }
    }
}

</script>

<style scoped>
</style>

