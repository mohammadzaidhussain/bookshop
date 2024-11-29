<template>
  <ion-app>
    <ion-router-outlet />
    <app-toast />
  </ion-app>
</template>

<script>
import { IonApp, IonRouterOutlet } from '@ionic/vue';
import storageService from './service/local_storage_service';
import AppToast from './components/AppToast.vue';
import {
  mapMutations
} from 'vuex';


export default {
  name: 'App',
  components: {
    IonApp,
    IonRouterOutlet,
    AppToast
  },
  methods: {
    ...mapMutations('authStore', {
      setLoggedInUser: 'SET_LOGGED_IN_USER'
    })
  },
  mounted() {
    const userData = storageService.getPersistedData("logged_in_user");
    this.setLoggedInUser(userData);
  }
}
</script>
