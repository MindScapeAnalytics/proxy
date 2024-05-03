#!/bin/bash

# Адрес Docker Registry, который мы хотим добавить в список небезопасных регистров
DOCKER_REGISTRY="dev.docker.registry.msanalytics.ru:4443"

# Путь к файлу конфигурации Docker Daemon
DOCKER_CONFIG="/etc/docker/daemon.json"

# Флаг для отслеживания изменений в конфигурации
CHANGED=false

# Проверка существования файла конфигурации
if [ -f "$DOCKER_CONFIG" ]; then
  # Если файл существует, проверим, уже ли добавлен наш Docker Registry
  if ! grep -q "$DOCKER_REGISTRY" "$DOCKER_CONFIG"; then
    # Если нет, добавляем в список insecure-registries
    jq --arg registry "$DOCKER_REGISTRY" \
      '.["insecure-registries"] += [$registry]' "$DOCKER_CONFIG" > temp.json \
      && mv temp.json "$DOCKER_CONFIG"
    CHANGED=true
  fi
else
  # Если файла нет, создаем его с нужной конфигурацией
  echo "{ \"insecure-registries\": [\"$DOCKER_REGISTRY\"] }" > "$DOCKER_CONFIG"
  CHANGED=true
fi

# Перезапускаем Docker только если конфигурация изменилась
if [ "$CHANGED" = true ]; then
  echo "Перезапускаем Docker для применения изменений..."
  sudo systemctl restart docker
  echo "Docker перезапущен. Docker Registry добавлен в список небезопасных регистров."
else
  echo "Docker Registry уже в списке небезопасных регистров. Перезапуск Docker не требуется."
fi
