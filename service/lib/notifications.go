package lib

/*
my sub map_notification_actions {
    my @actions = ();
    while (@_) {
        my $identifier = shift;
        my $localized  = shift;
        my @pair       = ( $identifier, $localized );
        push( @actions, \@pair );
    }
    return \@actions;
}

my sub map_notifications {

    # {
    #   "type":"method_call",
    #   "endian":"l",
    #   "flags":0,
    #   "version":1,
    #   "cookie":9,
    #   "timestamp-realtime":1744565025716135,
    #   "sender":":1.135",
    #   "destination":":1.45",
    #   "path":"/org/freedesktop/Notifications",
    #   "interface":"org.freedesktop.Notifications",
    #   "member":"Notify",
    #   "payload":{
    #     "type":"susssasa{sv}i",
    #     "data":[
    #       "notify-send",
    #       0,
    #       "",
    #       "Party time!",
    #       "It is time to party hard",
    #       ["dance", "Dance party!"],
    #       {
    #         "urgency":{
    #           "type":"y",
    #           "data":1
    #         },
    #         "sender-pid":{
    #           "type":"x"
    #           "data":129140
    #         }
    #       },
    #       -1
    #     ]
    #   }
    # }
    my $method_call = decode_json $_;
    my $payload     = $method_call->{'payload'};
    my $data        = $payload->{'data'};

    my $app_name       = $data->[0];
    my $replaces_id    = $data->[1];
    my $app_icon       = $data->[2];
    my $summary        = $data->[3];
    my $body           = $data->[4];
    my $actions        = &map_notification_actions( $data->[5] );
    my $hints          = $data->[6];
    my $expire_timeout = $data->[7];

    # TODO: Add timestamp
    my %event = (
        type           => 'Notification',
        app_name       => $app_name,
        replaces_id    => $replaces_id,
        app_icon       => $app_icon,
        summary        => $summary,
        body           => $body,
        actions        => $actions,
        hints          => $hints,
        expire_timeout => $expire_timeout
    );

    return \%event;
}

my sub notifications_worker {
    my @command = (
        'busctl', 'monitor', '--user',
        '--destination=org.freedesktop.Notifications',
        "--match=member='Notify'", '--json=short'
    );

    &listen_process( 'notifications', \@command, \&map_notifications );

    return;
}
*/
