<template>
    <Teleport to="body">
        <ion-toast
            :is-open="getToastState?.visible"
            :message="getToastState?.message"
            :duration="getToastState?.defaultTimeoutInMs"
            :position="getToastState?.position"
            position-anchor="header"
            :button="toastButtons"
            :class="['toast', getToastState?.type || 'success']"
            @didDismiss="toggle(false)"
        >
            <!-- <template #message>
                <div class="toast-content">
                    <span>{{ getToastState?.message }}</span>
                    <ion-icon
                        name="close-circle"
                        class="close-button"
                        @click="toggle(false)"
                    >
                    </ion-icon>
                </div>
            </template> -->
            
        </ion-toast>
    </Teleport>
</template>

<script>

import { IonIcon, IonToast } from '@ionic/vue';
import { closeOutline } from 'ionicons/icons';
import {
    mapGetters,
    mapActions
} from 'vuex';

export default {
    name: 'AppToast',
    components: {
        IonToast,
        IonIcon
    },
    computed: {
        ...mapGetters('commonStore', {
            getToastState: 'getToastState'
        }),
    },
    data() {
        return {
            closeOutline,
            toastButtons: [
                {
                    text: 'Dismiss',
                    role: 'cancel',
                },
            ]
        }
    },
    methods: {
        ...mapActions('commonStore', {
            toggleToast: 'toggleToast'
        }),
        toggle(visibility, message = '', type = '') {
            this.toggleToast({
                visible: visibility,
                message: message,
                type: type,
                defaultTimeoutInMs: 5000
            })
        }
    }
}
</script>


<style scoped>
.toast {
    --background: transparent;
    --color: white;
    position: fixed;
    top: 1rem;
    right: 1rem;
    z-index: 9999;
    display: inline-block;
}

.toast-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.close-button {
    margin-left: 1rem;
    cursor: pointer;
    font-size: 1.5rem;
    color: inherit;
}

ion-toast.success {
    --background: green !important;
    --color: white !important;
}

ion-toast.error {
    --background: red !important;
    --color: white !important;
}


ion-toast.warning {
    --background: rgb(214, 214, 1) !important;
    --color: white !important;
}

ion-toast.info {
    --background: rgb(0, 119, 255) !important;
    --color: white !important;
}



</style>