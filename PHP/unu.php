<?php

class UNURequest
{
    private static $apiURL = 'https://u.nu/api.php';
    private $url;
    private $keyword;
    private $title;
    private $username;
    private $password;

    /**
     * @return string
     */
    public function getURL()
    {
        return $this->url;
    }

    /**
     * @param string $url
     * @return $this
     */
    public function setURL($url)
    {
        $this->url = $url;
        return $this;
    }

    /**
     * @return string
     */
    public function getKeyword()
    {
        return $this->keyword;
    }

    /**
     * @param string $keyword
     * @return $this
     */
    public function setKeyword($keyword)
    {
        $this->keyword = $keyword;
        return $this;
    }

    /**
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * @param string $title
     * @return $this
     */
    public function setTitle($title)
    {
        $this->title = $title;
        return $this;
    }

    /**
     * @return string
     */
    public function getUsername()
    {
        return $this->username;
    }

    /**
     * @param string $username
     * @return $this
     */
    public function setUsername($username)
    {
        $this->username = $username;
        return $this;
    }

    /**
     * @return string
     */
    public function getPassword()
    {
        return $this->password;
    }

    /**
     * @param string $password
     * @return $this
     */
    public function setPassword($password)
    {
        $this->password = $password;
        return $this;
    }

    private function buildQuery()
    {
        $payload = array();
        $payload['url'] = $this->url;
        if (!empty($this->title)) {
            $payload['title'] = $this->title;
        }
        if (!empty($this->keyword)) {
            $payload['keyword'] = $this->keyword;
        }
        if (!empty($this->username)) {
            $payload['username'] = $this->username;
        }
        if (!empty($this->password)) {
            $payload['password'] = $this->password;
        }
        $payload['action'] = 'shorturl';
        $payload['format'] = 'json';
        return http_build_query($payload);
    }

    public function submit()
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, static::$apiURL);
        curl_setopt($ch, CURLOPT_HEADER, 0);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $this->buildQuery());
        $data = curl_exec($ch);
        curl_close($ch);
        var_dump($data);
        return json_decode($data);
    }
}