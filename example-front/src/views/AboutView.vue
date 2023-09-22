<template>
  <div class="about">
    <section v-if="errored">
      <p>Nous sommes désolés, nous ne sommes pas en mesure de récupérer ces informations pour le moment. Veuillez réessayer ultérieurement.</p>
      {{ errorMessage }}
    </section>

    <section v-else>
      <div v-if="loading">Chargement...</div>

      <div
        v-else
        v-for="value in info"
      >
        Nom: {{ value["name"] }} <br>
        Durée: {{ value["time"] }} <br>
        Description: {{ value["summary"] }} <br>
        -------
      </div>

    </section>
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>

<script lang="ts">
import axios from 'axios';

import { defineComponent } from "vue";

export default defineComponent({
  name: "tutorials-list",
  data () {
    return {
      info: [],
      loading: true,
      errored: false,
      errorMessage: ""
    }
  },
  mounted () {
    axios
      .get("http://localhost:4000/cours")
      .then(response => {
        this.info = response.data
      })
      .catch(error => {
        console.log(error)
        this.errored = true
        this.errorMessage = error
      })
      .finally(() => this.loading = false)
  }
});

</script>