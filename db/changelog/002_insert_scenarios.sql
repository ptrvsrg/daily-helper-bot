INSERT INTO scenarios(name)
VALUES ('start_not_authorized'),
       ('start_authorized'),
       ('profile'),
       ('projects'),
       ('project'),
       ('tasks'),
       ('task')
ON CONFLICT DO NOTHING;
