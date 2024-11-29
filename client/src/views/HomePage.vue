<template>
  <ion-menu content-id="main-content" ref="menu">
    <ion-header>
      <ion-toolbar>
        <ion-title>
            Menu
        </ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content>
      <ion-list>
        <ion-item router-link="/login" v-if="!getLoggedInUser">
          <app-label>
             <p>
                {{ loginText }}
             </p>
          </app-label>
        </ion-item>
        <div  v-else>
          <ion-item router-link="/books">
            <app-label>
              <p>
                  {{ bookList }}
              </p>
            </app-label>
          </ion-item>
          <ion-item router-link="/staffs" v-if="getLoggedInUser?.role != 'staff'">
            <app-label>
              <p>
                  {{ staffListText }}
              </p>
            </app-label>
          </ion-item>
          <ion-item @click="logout">
            <app-label>
              <p>
                  {{ logoutText }}
              </p>
            </app-label>
          </ion-item>
        </div>
      </ion-list>
    </ion-content>
  </ion-menu>
  <ion-page id="main-content">
    <ion-header id="header">
      <ion-toolbar>
        <ion-buttons slot="start">
          <ion-menu-button></ion-menu-button>
        </ion-buttons>
        <ion-title>{{ appName }}</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content class="ion-padding pt-20">
      <ion-router-outlet />
    </ion-content>
  </ion-page>
</template>

<script>
import { IonButtons, IonContent, IonHeader, IonItem, IonList, IonMenu, IonMenuButton, IonPage, IonRouterOutlet, IonTitle, IonToolbar } from '@ionic/vue';
import { AppName, Login, BookList, LogoutText, StaffListText } from '../constants/constant.js';
import AppLabel from '../components/AppLabel.vue';
import {
  mapGetters,
  mapMutations
} from 'vuex';

export default {
  name: 'HomePage',
  components: {
    IonMenu,
    IonContent, 
    IonHeader, 
    IonPage, 
    IonTitle, 
    IonToolbar,
    IonList,
    IonItem,
    AppLabel,
    IonRouterOutlet,
    IonButtons,
    IonMenuButton
  },
  data() {
    return {
      appName: AppName,
      loginText: Login,
      bookList: BookList,
      logoutText: LogoutText,
      staffListText: StaffListText
    }
  },
  computed: {
    ...mapGetters('authStore', {
      getLoggedInUser: 'getLoggedInUser'
    })
  },
  methods: {
    ...mapMutations('authStore', {
      clearAuthStore: 'CLEAR_DATA'
    }),
    async logout() {
      this.clearAuthStore();
      this.$router.push('/')
    }
  },
}
</script>

<style scoped>
#container {
  text-align: center;
  
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
}

#container strong {
  font-size: 20px;
  line-height: 26px;
}

#container p {
  font-size: 16px;
  line-height: 22px;
  
  color: #8c8c8c;
  
  margin: 0;
}

#container a {
  text-decoration: none;
}
</style>
