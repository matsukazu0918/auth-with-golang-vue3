<template>
  <nav class="navbar navbar-expand-md navbar-dark bg-dark">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item">
        <router-link to="/" class="nav-link">Home</router-link>
      </li>
    </ul>

    <ul class="navbar-nav my-2 my-lg-0">
      <!-- ログイン済みなら表示 -->
      <template v-if="auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link" @click="logout"
            >Logout</router-link
          >
        </li>
      </template>
      <!-- 未ログインなら表示 -->
      <template v-if="!auth">
        <li class="nav-item">
          <router-link to="/login" class="nav-link">Login</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/register" class="nav-link">Register</router-link>
        </li>
      </template>
    </ul>
  </nav>
</template>

<script>
import { computed } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import axios from "axios";

export default {
  name: "Nav",
  setup() {
    const store = useStore();
    const router = useRouter();
    const auth = computed(() => store.state.auth);
    const logout = async () => {
      await axios.post("logout", {});
      store.dispatch("setaAuth", false);
      await router.push("/login");
    };
    return {
      auth,
    };
  },
};
</script>