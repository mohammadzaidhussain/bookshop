<template>
    <ion-toggle :checked="checked" :disabled="disabled" :enable-on-off-label="enableOnOffLabel" @click="(event) => this.$emit('change', event)">
        {{ title }}
    </ion-toggle>
</template>


<script>

import { IonToggle } from '@ionic/vue';

export default {
    name: 'AppSwitch',
    components: {
        IonToggle
    },
    props: {
        title: {
            type: String,
            default: ''
        },
        enableOnOffLabel: {
            type: Boolean,
            default: true
        },
        checked: {
            type: Boolean,
            default: true
        },
        disabled: {
            type: Boolean,
            default: false
        },
    }
}
</script>