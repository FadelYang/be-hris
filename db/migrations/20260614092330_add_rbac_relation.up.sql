CREATE TABLE users_roles (
  user_id UUID NOT NULL,
  role_id UUID NOT NULL,

  PRIMARY KEY (user_id, role_id),

  CONSTRAINT fk_users_roles_user
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_users_roles_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE CASCADE
);

CREATE TABLE roles_menus_permissions (
  role_id UUID NOT NULL,
  menu_id UUID NOT NULL,
  permission_id UUID NOT NULL,

  PRIMARY KEY (role_id, menu_id, permission_id),

  CONSTRAINT fk_roles_menus_permissions_role
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_roles_menus_permissions_menu
    FOREIGN KEY (menu_id)
    REFERENCES menus(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_roles_menus_permissions_permission
    FOREIGN KEY (permission_id)
    REFERENCES permissions(id)
    ON DELETE CASCADE
);